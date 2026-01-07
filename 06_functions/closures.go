package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("counter", counter)
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
