package main

func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
            close channel here to avoid deadlock
        }
    }()

 when count finishes no goroutines have access to ch so deadlock
    for n := range ch {
        println(n)
    }
}


0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox3288101957/prog.go:11 +0xa8
