package main

import "fmt"

func main() {
	type NoKTP string

	var ktpUdin NoKTP = "0989123812"
	fmt.Println(ktpUdin)
}
