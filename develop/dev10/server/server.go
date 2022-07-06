package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	//Listening to tcp at port
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	//Accepting connections
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//On \n parsing command
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		//If CLOSE command comes, closing server
		if strings.TrimSpace(string(netData)) == "CLOSE" {
			fmt.Println("Exiting TCP server!")
			c.Close()
			return
		}

		//Print message accepted
		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"

		//Write back current time
		_, err = c.Write([]byte(myTime))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
