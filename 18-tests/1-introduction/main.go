package main

import (
	"fmt"

	"github.com/golang-course/18-tests/1-introduction/address"
)

func main() {
	adressType := address.AddressType("avenue")
	fmt.Println(adressType)
}
