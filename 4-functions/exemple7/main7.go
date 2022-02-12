package main

// closure function
// it save the variable when it was declared in the function

import "fmt"

func closure() func() {
	text := "Inside closure function"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Inside main function"
	fmt.Println(text)

	newFunction := closure()
	newFunction()

}
