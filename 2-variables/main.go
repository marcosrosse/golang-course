package main

import "fmt"

func main() {
	var variable1 string = "Variable 1"
	variable2 := "Variable 2"

	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "Variable 3"
		variable4 string = "Variable 4"
	)

	fmt.Println(variable3)
	fmt.Println(variable4)

	variable5, variable6 := "Variable 5", "Variable 6"

	fmt.Println(variable5)
	fmt.Println(variable6)

	fmt.Println(variable1, variable2, variable3, variable4, variable5, variable6)

	const constant1 string = "Constant 1"
	fmt.Println(constant1)

	// inverse of var
	variable1, variable2 = variable2, variable1
	fmt.Println(variable1, variable2)
}
