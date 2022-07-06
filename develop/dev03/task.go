package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Creating flags
var (
	ReverseFlag    = flag.Bool("r", false, "use to reverse order")
	UniqueFlag     = flag.Bool("u", false, "use only to list unique lines")
	NumberSortFlag = flag.Bool("n", false, "use to sort numbers")
	ColumnSortFlag = flag.Int("k", -1, "use to sort columns")
)

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

//Function to write lines into a file
func writeLines(file string, lines []string) error {
	//Creating a file and closing it after writing into it
	f, err := os.Create("sorted_" + file)
	if err != nil {
		return err
	}
	defer f.Close()

	//Creating new writer
	w := bufio.NewWriter(f)
	defer w.Flush()

	//If reverse write is invoked rearranging slice
	if *ReverseFlag {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	//If unique flag is invoked deleting repeating values
	if *UniqueFlag {
		inResult := make(map[string]bool)
		var result []string
		for _, str := range lines {
			if _, ok := inResult[str]; !ok {
				inResult[str] = true
				result = append(result, str)
			}
		}
		lines = result
	}

	//Writing strings into file
	for _, line := range lines {
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

//Function to sort numbers is NumFlag is invoked
func numSort(lines []string) ([]string, error) {
	var a []int
	var s []string
	for _, v := range lines {
		int, err := strconv.Atoi(strings.TrimSuffix(v, "\n"))
		if err != nil {
			return nil, err
		}
		a = append(a, int)
	}
	sort.Ints(a)
	for _, v := range a {
		s = append(s, strconv.Itoa(v)+"\n")
	}
	lines = s

	return lines, nil
}

//Function to sort columns if ColumnFlag is invoked
func columnSort(lines []string, column int) []string {
	//Splitting lines into columns
	var colToSort [][]string
	for _, v := range lines {
		curLine := strings.Split(v, " ")
		//check if column to sort is within range
		if len(curLine)-1 < column {
			fmt.Fprintf(os.Stderr, "Invalid column: %d\n", column)
			os.Exit(1)
		}
		colToSort = append(colToSort, curLine)
	}
	//Sorting slice by increasing depending on column selected
	sort.Slice(colToSort, func(i, j int) bool {
		// edge cases
		if len(colToSort[i]) == 0 && len(colToSort[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(colToSort[i]) == 0 || len(colToSort[j]) == 0 {
			return len(colToSort[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}
		return colToSort[i][column] < colToSort[j][column]
	})
	var newLines []string
	for _, v := range colToSort {
		newLines = append(newLines, strings.Join(v, " "))
	}
	lines = newLines

	return lines
}

func main() {
	//Parsing flags
	flag.Parse()

	//Setting file to sort
	file := `columns.txt`

	//Reading lines
	lines, err := readLines(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading lines: %v\n", err)
		os.Exit(1)
	}

	//If numbers or columns are involved sort numbers or columns else sort strings
	if *NumberSortFlag {
		lines, err = numSort(lines)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sorting nums: %v\n", err)
			os.Exit(1)
		}
	} else if *ColumnSortFlag >= 0 {
		lines = columnSort(lines, *ColumnSortFlag)
	} else {
		sort.Strings(lines)
	}

	//Write lines to output file
	err = writeLines(file, lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
