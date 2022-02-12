package main

import "fmt"

// variatic functions

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func writeConsole(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	sumTotal := sum(1, 1)
	fmt.Println(sumTotal)

	writeConsole("Macarrone", 1, 2, 3, 4)
}
