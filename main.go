package main

import (
	"fmt"
	f "go-practice/functions"
)

func main() {
	a, p := f.Rectangle(100, 50)
	fmt.Println("Area: ", a)
	fmt.Println("Parameter: ", p)

	var age = 27
	var name = "John"

	fmt.Println("Before: ", name, age)
	f.PassingAddress(&age, &name) // `&` to get the address

	fmt.Println("After: ", name, age)

	f.AnonymousFunc()
}
