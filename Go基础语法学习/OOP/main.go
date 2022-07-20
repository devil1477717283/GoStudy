package main

import (
	"OOP/singeleton"
	"fmt"
)

func main() {
	newUser := User.GetUserInstance()
	fmt.Printf("%p", newUser)
}
