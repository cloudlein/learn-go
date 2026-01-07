package main

import "fmt"

func main() {
	//counter := 1

	//for counter <= 10 {
	//	fmt.Println("counter is ", counter)
	//	counter++
	//}

	for counter := 1; counter <= 20; counter++ {
		fmt.Println("counter is ", counter)
	}

	num := []int{1, 2, 3, 4, 5}

	for index, num := range num {
		fmt.Println(index, num)
	}

	// jika tidak butuh index
	for _, num := range num {
		fmt.Println(num)
	}

}
