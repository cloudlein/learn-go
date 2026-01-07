package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Udin", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("shit", filter)
}

func spamFilterTest(name string) string {
	if name == "Shit" {
		return "..."
	} else {
		return name
	}
}

func sayHelloWithFilterTest(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}
