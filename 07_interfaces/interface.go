package main

import "fmt"

func main() {
	fmt.Println()
}

type HasName interface {
	GetName() string
}

func sayHello(name HasName) {
	fmt.Println(name.GetName())
}
