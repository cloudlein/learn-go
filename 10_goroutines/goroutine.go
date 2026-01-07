package main

import(
	"fmt"
	"time"
)


func printMessage(message string) {
	fmt.Println(message)
}

func main(){
	go printMessage("Hello from goroutine!")
	time.Sleep(1 * time.Second) // Wait for goroutine to finish
}