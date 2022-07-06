package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	FieldFlag     = flag.Int("f", 0, "choose a specific field (column)")
	DelimFlag     = flag.String("d", "\t", "set dilimeter")
	SeparatedFlag = flag.Bool("s", false, "print only lines containing delimeters")
)

func main() {
	flag.Parse()
	//If app used without specification
	if *FieldFlag <= 0 && !*SeparatedFlag {
		fmt.Println("Field Not Specified, Exiting..")
		os.Exit(0)
	}

	//geting file name
	file := flag.Arg(0)

	//reading lines from file
	linesIn, err := readLines(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading lines: %v\n", err)
		os.Exit(1)
	}

	linesOut := cut(linesIn)

	for _, v := range linesOut {
		fmt.Println(v)
	}

}

func cut(linesIn []string) []string {
	var linesOut []string

	//If flags are set
	if *FieldFlag > 0 && !*SeparatedFlag {
		for _, v := range linesIn {
			//if using default delimiter TAB
			if *DelimFlag == "\t" {
				s := strings.Fields(v)
				if len(s) >= *FieldFlag {
					linesOut = append(linesOut, strings.TrimSuffix(s[*FieldFlag-1], "\n"))
				}
				//Using any other delimiter
			} else {
				s := bytes.Split([]byte(v), []byte(*DelimFlag))
				if len(s) >= *FieldFlag && len(s) > 1 {
					linesOut = append(linesOut, strings.TrimSuffix(string(s[*FieldFlag-1]), "\n"))
				}
			}
		}
	}

	if *SeparatedFlag {
		for _, v := range linesIn {
			//Check if line containes a delimiter (if it is splits into fields)
			if *DelimFlag == "\t" {
				s := strings.Fields(v)
				if len(s) > 1 {
					linesOut = append(linesOut, strings.TrimSuffix(v, "\n"))
				}
			} else {
				s := bytes.Split([]byte(v), []byte(*DelimFlag))
				if len(s) > 1 {
					linesOut = append(linesOut, strings.TrimSuffix(v, "\n"))
				}
			}
		}
	}

	return linesOut
}

//Function to read all lines from file and put them into a slice of strings
func readLines(file string) ([]string, error) {
	var lines []string
	//Open and defer close file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//New reader from file
	r := bufio.NewReader(f)
	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			//In case there is no delimiter in line
			if err != nil {
				line += string(delim)
			}
			//Add line to the slice
			lines = append(lines, line)
		}
		//If line is empty and there is an error
		if err != nil {
			//If EOF is reached
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return lines, nil
}
