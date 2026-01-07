package main

import "fmt"

type User struct {
    Name string
}

func changeName(u *User) {
    u.Name = "Budi"
}

func main() {
    user := User{Name: "Andi"}
    changeName(&user)
    fmt.Println(user.Name) // Budi
}
