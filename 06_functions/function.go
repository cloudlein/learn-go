package main

import "fmt"

func main() {
	sayHello()
	sayHelloParam("udin", "salahudin")
	fmt.Println(getHello("salahudin"))
	fmt.Println(getFullName("Salahudin", "Rina"))
	fmt.Println(getCompletedName())
}

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloParam(firstName string, lastName string) {
	fmt.Println("Hello " + firstName + " " + lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

// return multiple
func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

// return named value
func getCompletedName() (firstName, lastName string) {
	firstName, lastName = "udin", "rin"
	return firstName, lastName
}
