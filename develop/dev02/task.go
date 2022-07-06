package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ASCII for single digit numbers
const (
	zero rune = 48
	nine rune = 57
)

func main() {
	//Creating reader from stdIn
	r := bufio.NewReader(os.Stdin)
	//Reading line before \n
	line, _, err := r.ReadLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading line: %v\n", err)
		os.Exit(1)
	}

	//Unpacking string
	strOut, err := unpackStr(string(line))
	if err != nil || len(strOut) == 0 {
		fmt.Fprintf(os.Stderr, "error unpacking string: %s\n", string(line))
		os.Exit(1)
	}

}

func unpackStr(strIn string) (string, error) {
	//Creating strings Builder
	var b strings.Builder

	for i, v := range []rune(strIn) {
		//Checking if v is a number between 0 and 9
		if v >= zero && v <= nine {

			//Checking if the string doesnt start with a number or is there is a number before or after the rune being processed
			if i != 0 && ([]rune(strIn)[i-1] < zero || []rune(strIn)[i-1] > nine) {
				//Conversting number to int
				int, err := strconv.Atoi(string(v))
				if err != nil {
					return "", err
				}
				//adding neccessary number of a letter previous to the number being processed
				for ii := 1; ii < int; ii++ {
					b.WriteRune([]rune(strIn)[i-1])
				}
			} else {
				return "", nil
			}

		} else {
			//Printing not a number
			b.WriteRune(v)
		}
	}
	return b.String(), nil

}
