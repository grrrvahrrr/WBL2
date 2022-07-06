package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Output to set file name
var (
	FilenameFlag = flag.String("o", "", "use to specify output filename")
)

func Wget(url, fileName string) {

	resp := getResponse(url)
	writeToFile(fileName, resp)
}

// Make the GET request to a url, return the response
func getResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listening to keyboard: %v\n", err)
		os.Exit(1)
	}
	return resp
}

// Write the response of the GET request to file
func writeToFile(fileName string, resp *http.Response) {
	// Creating file to write to
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	//Creating writer
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating writer: %v\n", err)
		os.Exit(1)
	}

	//Copy response body to file
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error copying t buffer: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if *FilenameFlag == "" {
		Wget(flag.Arg(0), "file_downloaded")
	} else {
		Wget(flag.Arg(0), *FilenameFlag)
	}

}
