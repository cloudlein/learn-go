package main

import "fmt"

func main() {
	name := "salahudin"

	switch name {
	case "budi":
		fmt.Println("I'm budi")
	case "salahudin":
		fmt.Println("I'm salahudin")
	default:
		fmt.Println("name not recognized")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("name is too long")
	//case false:
	//	fmt.Println("name is too short")
	//default:
	//	fmt.Println("name true")
	//}

	length := len(name)
	switch {
	case length < 5:
		fmt.Println("name is too short")
	case length > 6:
		fmt.Println("name is too long")
	default:
		fmt.Println("name is recognized")
	}
}
