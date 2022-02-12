package main

import "fmt"

// SImple functions

func soma(a, b int) int {
	return a + b
}

func mathsCalc(n1, n2 int8) (int8, int8) {
	return n1 + n2, n1 - n2
	// sum := n1 + n2
	// sub := n1 - n2
}

func main() {
	sum := soma(2, 2)
	fmt.Println(sum)

	// variable as a function

	var c = func() {
		fmt.Println("Hello i am a function as variable")
	}
	c()

	//------------------------------------------------------

	var d = func(txt string) {
		fmt.Println(txt)
	}
	d("Hello i am a function as variable using parameter")

	//------------------------------------------------------

	var e = func(txt string) string {
		return txt
	}
	fmt.Println(e("Hello i am a function as variable using parameter and returning a string"))

	//------------------------------------------------------

	// function returning two values
	sumCalc, subCalc := mathsCalc(2, 100)
	fmt.Println(sumCalc, subCalc)

	// sumCalc, _ := mathsCalc(2, 100) // _ is a blank identifier to ignore the second value
	// fmt.Println(sumCalc)

}
