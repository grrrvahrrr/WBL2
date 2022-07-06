package main

import (
    "fmt"
)

func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}

A defer statements pushes a function call onto a stack. The stack of saved calls popped out (LIFO) and deferred functions are invoked immediately before the surrounding function returns.

After 1 is returned, the defer func() { i++ }() gets executed. Hence, in order of executions:

x = 1 (return 1)
x++ (defer func pop out from stack and executed)
x == 2 (final result of named variable i)


func anotherTest() int {
    var x int
    defer func() {
        x++
        add fmt.Println(x) to print 2
    }()
    x = 1
    return x
}

Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked.


func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}


Output:
2
1
