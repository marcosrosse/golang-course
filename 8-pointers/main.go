package main

import "fmt"

func main() {

	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable1++

	fmt.Println(variable1, variable2)

	// Pointers

	var variable3 int = 10
	var pointer *int = &variable3

	fmt.Println(variable3, pointer)  // show the memory address
	fmt.Println(variable3, *pointer) // show the value of the pointer

}
