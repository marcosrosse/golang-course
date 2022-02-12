package main

// function  panic and recover

import "fmt"

func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Recover completed!")
	}
}

func approvedStudent(n1, n2 float32) bool {

	defer recoverExec()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("The grade is exactly 6!")
}

func main() {
	fmt.Println(approvedStudent(6, 6))

}
