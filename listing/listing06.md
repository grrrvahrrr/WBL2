package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
    when append we create a new slice named i
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
    function doesnt return the new overwritten slice
}


[3 2 3]
