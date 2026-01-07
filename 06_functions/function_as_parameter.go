package main

import "fmt"

func main() {
	sayHelloWithFilter("Udin", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("shit", filter)
}

func spamFilter(name string) string {
	if name == "Shit" {
		return "..."
	} else {
		return name
	}
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}
