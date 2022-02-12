package main

import "fmt"

// functions and pointers

func invertSignal(number int) int {
	return number * -1
}

func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	invertedNumber := invertSignal(number)
	fmt.Println(invertedNumber)

	newNumber := 80
	fmt.Println(newNumber)

	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}
