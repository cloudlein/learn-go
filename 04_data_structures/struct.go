package main

import "fmt"

func main() {
	var data Customer
	data.Name = "Salahudin"
	data.Address = "Indonesia"
	data.Age = 22

	data_two := Customer{
		Name:    "udin",
		Address: "Indonesia",
		Age:     18,
	}

	data_three := Customer{"udin", "Indonesia", 30}

	data_four := Biodata{
		Name: "udin",
	}

	fmt.Println(data_four.sayHello)

	//fmt.Println(data)
	//fmt.Println(data.Name)
	//fmt.Println(data.Address)
	//fmt.Println(data.Age)
	//
	//fmt.Println(data_two)
	//fmt.Println(data_three)
}

type Customer struct {
	Name, Address string
	Age           int
}

func (bio Biodata) sayHello() {
	fmt.Printf("Hello, %s!\n", bio.Name)
}

type Biodata struct {
	Name string
}
