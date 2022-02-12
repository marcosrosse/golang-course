package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64 := int8(127), int16(32767), int32(2147483647), int64(9223372036854775807)
	var number1 int = 1000
	fmt.Println(number1)

	var number2 uint = 127 // only positive
	fmt.Println(number2)

	//alias for int32
	var number3 rune = 123456
	fmt.Println(number3)

	// alias for uint8
	var number4 byte = 255
	fmt.Println(number4)

	//------------------------------------------------------
	// float32, float64
	var number5 float32 = 1.2345
	fmt.Println(number5)
	var number6 float64 = 198439489384.2345
	fmt.Println(number6)

	// cannot declare float only like int, you need to specify the type 32 or 64

	//------------------------------------------------------

	// char

	var str string = "Hello World"
	fmt.Println(str)

	str2 := "Hello World string 2"
	fmt.Println(str2)

	char := 'a'
	fmt.Println(char) // char print as int

	//------------------------------------------------------

	// bool
	var bool1 bool = true
	fmt.Println(bool1)
	var bool2 bool = false
	fmt.Println(bool2)

	//------------------------------------------------------

	// error

	var err error
	fmt.Println(err)

	var err2 error = errors.New("Internal server error")
	fmt.Println(err2)

}
