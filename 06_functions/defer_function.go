package main

import "fmt"

func main() {
	runApplication()
}

func logging() {
	fmt.Println("done run application")
}

func runApplication() {
	defer logging()
	fmt.Println("Running application Go!")
}
