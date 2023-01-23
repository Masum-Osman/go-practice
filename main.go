package main

import (
	"fmt"
	f "go-practice/functions"
)

func main() {
	a, p := f.Rectangle(100, 50)
	fmt.Println("Area: ", a)
	fmt.Println("Parameter: ", p)
}
