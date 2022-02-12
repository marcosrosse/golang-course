package main

// function defer

import "fmt"

// func func1() {
// 	fmt.Println("Executing func 1")
// }

// func func2() {
// 	fmt.Println("Executing func 2")
// }

func approvedStudent(n1, n2 float32) bool {
	defer fmt.Println("Calculated average. The result will be returned")
	fmt.Println("Entering the function to check if the student is approved")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(approvedStudent(6, 3))

	// defer func1()
	// func2()
}
