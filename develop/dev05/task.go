package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	AfterFlag      = flag.Int("A", 0, "use to print lines after match")
	BeforeFlag     = flag.Int("B", 0, "use to print lines before match")
	ContextFlag    = flag.Int("C", 0, "use to print lines around match")
	CountFlag      = flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file.  With the -v, --invert-match option (see below), count non-matching lines.")
	IgnoreCaseFlag = flag.Bool("i", false, "Ignore case distinctions in patterns and input data, so that characters that differ only in case match each other.")
	InvertFlag     = flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines.")
	FixedFlag      = flag.Bool("F", false, "Interpret PATTERNS as fixed strings, not regular expressions.")
	LineNumFlag    = flag.Bool("n", false, "Prefix each line of output with the 1-based line number within its input file.")
)

func main() {
	flag.Parse()

	result, count, err := grep("file.txt", "line")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error grep: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
	fmt.Println(count)

}

func grep(file string, pattern string) ([]string, int, error) {
	//Setting context range
	if *ContextFlag > 0 {
		*AfterFlag = *ContextFlag
		*BeforeFlag = *ContextFlag
	}

	var result []string
	var count int
	//reading lines from file
	lines, err := readLines(file)
	if err != nil {
		return nil, 0, err
	}

	//Creating regular expretions object
	var r *regexp.Regexp
	if *IgnoreCaseFlag {
		//Regular expresion for ignoring case
		r = regexp.MustCompile("(?i)" + pattern)
	} else {
		r = regexp.MustCompile(pattern)
	}

	for i, v := range lines {
		//if regualr expretions not used
		if *FixedFlag {
			if v == pattern+"\n" {
				//print line number before target string
				if *LineNumFlag {
					result = append(result, strconv.Itoa(i+1)+" "+v)
				} else {
					result = append(result, v)
				}
			}
			//If regular expresions are used
		} else {
			str := r.FindString(v)
			if str != "" && !*InvertFlag {
				//Appending all lines before the match and the match
				if *BeforeFlag > 0 {
					for lineBefore := *BeforeFlag; lineBefore >= 0; lineBefore-- {
						if i-lineBefore < 0 {
							continue
						}
						if *LineNumFlag {
							result = append(result, strconv.Itoa(i-lineBefore+1)+" "+lines[i-lineBefore])
						} else {
							result = append(result, lines[i-lineBefore])
						}
					}
				}

				//Appending all lines after the match
				if *AfterFlag > 0 {
					//Check if the match wasnt added because of context flag
					if *ContextFlag <= 0 {
						if *LineNumFlag {
							result = append(result, strconv.Itoa(i+1)+" "+v)
						} else {
							result = append(result, v)
						}
					}
					for lineAfter := 1; lineAfter <= *AfterFlag; lineAfter++ {
						//Check if lines don't go out of range
						if i+lineAfter > len(lines)-1 {
							break
						}
						if *LineNumFlag {
							result = append(result, strconv.Itoa(i+lineAfter+1)+" "+lines[i+lineAfter])
						} else {
							result = append(result, lines[i+lineAfter])
						}
					}
				}

				//If only match needs to be printed
				if *AfterFlag <= 0 && *BeforeFlag <= 0 {
					if *LineNumFlag {
						result = append(result, strconv.Itoa(i+1)+" "+v)
					} else {
						result = append(result, v)
					}
				}
				//append lines that don't match
			} else if str == "" && *InvertFlag {
				if *LineNumFlag {
					result = append(result, strconv.Itoa(i+1)+" "+v)
				} else {
					result = append(result, v)
				}
			}
		}
	}

	//If only count is needed
	if *CountFlag {
		return nil, len(result), nil
	}

	return result, count, nil
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
