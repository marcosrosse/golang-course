package main

import "fmt"

// functions named returned

func calcMaths(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := calcMaths(10, 20)
	fmt.Println(sum, sub)
}
