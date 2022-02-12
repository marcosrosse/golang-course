package main

import "fmt"

type user struct {
	name   string
	age    int8
	adress address
}

type address struct {
	street string
	number uint8
}

func main() {
	var u user
	u.name = "John"
	u.age = 30
	fmt.Println(u)

	fmt.Println("------------")

	u2 := user{name: "Bob", age: 20}
	fmt.Println(u2)

	fmt.Println("------------")

	// u3 := user{name: "Marina", age: 18, adress: address{street: "Main", number: 123}}
	// Nesteds structs
	u3 := user{
		name:   "Marina",
		age:    18,
		adress: address{"Odemis", 0},
	}
	fmt.Println(u3)
}
