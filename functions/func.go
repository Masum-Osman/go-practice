package functions

func Rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b) // see, no `:` in this line. as the var in previously descrived on return statement
	area = l * b

	return
}

func PassingAddress(a *int, t *string) {
	*a = *a + 5 // `*` to point the address
	*t = *t + " Doe"
}
