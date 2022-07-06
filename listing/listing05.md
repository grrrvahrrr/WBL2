package main

type customError struct {
     msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
     {
         // do something
     }
     return nil
}

func main() {
    var err error
    err = test()
    interface error is not nil because is has type of customError with nil value
    if err != nil {
        println("error")
        return
    }
    println("ok")
}


output:
error
