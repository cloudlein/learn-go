package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "salahudin",
		"age":  "23",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)
	fmt.Println(len(person))
	person["name"] = "Jack"
	fmt.Println(person["name"])
	fmt.Println(person)
	delete(person, "name")
	fmt.Println(person)
}
