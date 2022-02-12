package main

import "fmt"

func main() {
	fmt.Println("Array and Slice")

	// Arrays
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices
	slice := []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice)
	slice = append(slice, 21)
	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	array2[2] = "potato"
	fmt.Println(slice2)

	// an slice is a reference to an array
}
