package main

import "fmt"

func main() {
	// arithmetic operators
	sum := 2 + 2
	sub := 2 - 2
	div := 2 / 2
	mul := 2 * 2
	mod := 2 % 2

	fmt.Println(sum, sub, div, mul, mod)

	// i cannot use two vars with different types, like:
	// var number1 int8 = 2
	// var number2 int16 = 100
	// sum2 := number1 + number2
	// fmt.Fprintln(sum2)

	// assignment operators
	var var1 string = "Hello"
	var2 := "World"
	fmt.Println(var1, var2)

	// relational operators
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	fmt.Println(1 < 1)
	fmt.Println(1 > 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 >= 1)

	fmt.Println("------------------")

	// logical operators
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("------------------")

	// bitwise operators
	fmt.Println(1 & 1)

	fmt.Println("------------------")

	// unary operators
	number := 10
	number++
	number += 10
	fmt.Println(number)
	number--
	fmt.Println(number)
}
