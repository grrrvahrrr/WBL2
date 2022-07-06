package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}

in slice b there is a pointer to values 1 to 4 of a
[77 78 79]
