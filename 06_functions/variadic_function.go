package main

import "fmt"

func main() {
	total := sumAll(10, 10, 20, 100)
	fmt.Println(total)

	// mengirim data yang berbentuk slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total1 := sumAll(numbers...)
	fmt.Println(total1)

}

func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
