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
	f.ClosuresArea()
	f.ClosuresLoop()

	partial := f.HigherOrderSum(100)
	fmt.Println(partial(50))

	f.HigherOrderFuncAsParameter()
	squareSum := f.SquareSum(5)(6)(7)
	fmt.Println(squareSum)

	f.VariadicExample("red", "blue", "green", "yellow")
	fmt.Println(f.AreaCalculation("Square", 20))
	fmt.Println(f.AreaCalculation("Rectangle", 20, 30))
	f.VariadicDiffTypesOfParams(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}
