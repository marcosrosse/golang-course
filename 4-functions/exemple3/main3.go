package main

// anonymous functions

import "fmt"

func main() {
	func() {
		fmt.Println("Simple function without name")
	}()

	fmt.Println("------------------")

	func(text string) {
		fmt.Println(text)
	}("Function without name receiving parameter")

	fmt.Println("------------------")

	returnFunc := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Function without name receiving parameter and returning")
	fmt.Println(returnFunc)
}
