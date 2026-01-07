package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklisted.")
	} else {
		fmt.Println("Successfully registered user.")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "shit"
	}

	registerUser("John", blacklist)

	registerUser("shit", func(name string) bool {
		return name == "shit"
	})

}
