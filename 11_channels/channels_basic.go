package main

import "fmt"

func Worker(ch chan int) {
    ch <- 10
}


func main() {
	ch := make(chan string)
	ch_int := make(chan int)

	go func() {
		ch <- "Hello, Channels!"
	}()

	message := <-ch
	fmt.Println(message)
    go Worker(ch_int)

    result := <-ch_int
    fmt.Println(result)
}
