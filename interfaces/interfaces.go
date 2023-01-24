package interfaces

import "fmt"

type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id: ", e)
	fmt.Println("Employee name: ", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func RunInterface() {

	e1 := Emp(15)

	e1.PrintName("John Doe")
	fmt.Println("Employee Salary: ", e1.PrintSalary(25000, 5))
}
