package main

import "fmt"

func main() {
	number := 10

	if number > 15 {
		fmt.Println("The number is greater than 15")
	} else {
		fmt.Println("The number", number, "is lower than 15")
	}

	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("The number is greater than 0")
	} else {
		fmt.Println("The number", anotherNumber, "is lower than 0")
	}

}
