package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementing i: ", i)
		time.Sleep(time.Second)
	}

	fmt.Println("--------------------------------")

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementing j:", j)
		time.Sleep(time.Second)
	}
	fmt.Println("--------------------------------")

	names := [3]string{"John", "Maria", "Valeska"} // to transform in a slice only remove the number between []

	for i, name := range names {
		fmt.Println(i, name)
	}

	fmt.Println("--------------------------------")

	for i, letter := range "Pandão" {
		fmt.Println(i, string(letter)) //transforming in string
	}

	fmt.Println("--------------------------------")

	user := map[string]string{
		"name":     "Marcos",
		"nickname": "Rosse",
		"email":    "marcosrossem@gmail.com",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
	fmt.Println("--------------------------------")

}
