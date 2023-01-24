package functions

import (
	"fmt"
	"reflect"
)

// anonymous function:
var (
	area = func(l int, b int) int {
		return l * b
	}
)

func Rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b) // see, no `:` in this line. as the var in previously descrived on return statement
	area = l * b

	return
}

func PassingAddress(a *int, t *string) {
	*a = *a + 5 // `*` to point the address
	*t = *t + " Doe"
}

/*
	Anonymous Functions
*/

func AnonymousFunc() {
	fmt.Println(area(50, 50))

	fmt.Printf(
		"81 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(81),
	)
}

/*
	Closure Functions
*/

// Note: Closures are a special case of anonymous functions.
// Closures are anonymous functions which access the variables defined outside the body of the function.

func ClosuresArea() {
	l := 20
	b := 30

	func() {
		area := l * b
		fmt.Println("Closures Area: ", area)
	}()
}

func ClosuresLoop() {
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}

/*
	Higher Order Functions
*/

//Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.

//returning func()
func sum(x, y int) int {
	return x + y
}

func HigherOrderSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

//taking func() as argument
func twice(f func(int) int) func(int) int {
	return func(x int) int {
		return f(f(x))
	}
}

func HigherOrderFuncAsParameter() {
	plusThree := func(i int) int {
		return i + 3
	}

	g := twice(plusThree)

	fmt.Println(g(7))
}

// Returning function
func SquareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

/*
	Variadic Functions
*/
func VariadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

func AreaCalculation(str string, y ...int) int {
	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		}
		if str == "Square" {
			area = val * val
		}
	}

	return area
}

func VariadicDiffTypesOfParams(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, " -- ", reflect.ValueOf(v).Kind())
	}
}
