package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

here we return interface error of type os.PathError with nil value

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
    interface is only nil when its type is nil and its value is nil. fmt.Println(err == (*os.PathError)(nil)) is true
}

<nil>
false
