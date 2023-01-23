package functions

import "fmt"

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

func AnonymousFunc() {
	fmt.Println(area(50, 50))

	fmt.Printf(
		"81 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(81),
	)
}
