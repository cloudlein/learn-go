package main

import "fmt"

func main() {
	var name [2]string

	name[0] = "salahudin"
	name[1] = "muhammad salahudin"

	fmt.Println(name[0])
	fmt.Println(name[1])

	var values = [3]int{1, 2, 3}
	var values1 = [...]int{1, 2, 3}

	values[2] = 10
	fmt.Println(values)
	fmt.Println(values1)
	fmt.Println(values[2])
	fmt.Println(len(values))

}
