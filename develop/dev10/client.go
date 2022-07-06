package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

//Setup 10 sec timeout
var (
	TimeOutFlag = flag.Int("timeout", 10, "setup connection timeout")
)

func main() {
	flag.Parse()

	//Accept addres as first argumant
	CONNECT := flag.Arg(0)

	var timeout time.Duration = time.Duration(*TimeOutFlag) * time.Second

	//Connect to Server with timeout
	c, err := net.DialTimeout("tcp", CONNECT, timeout)
	if err != nil {
		fmt.Println(err)
		//Exit after timeout runs out
		<-time.After(timeout)
		os.Exit(1)
	}

	for {
		//Read message and print it
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		//Exit if recieved EOF
		if text == "" {
			fmt.Println("TCP client exiting...")
			return
		}

		//Print server response
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Connection Lost!")
			return
		}
		fmt.Print("->: " + message)
	}
}
