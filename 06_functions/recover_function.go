package main

import "fmt"

func main() {
	runApplicationTest(true)
}

func closeApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi error ", message)
}

// wrong way
//func runApplicationTest(error bool) {
//	defer endApp()
//
//	if error {
//		panic("error")
//	}
//
//	message := recover()
//	fmt.Println(message)
//}

func runApplicationTest(error bool) {
	defer closeApp()

	if error {
		panic("error")
	}

}
