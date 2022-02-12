package main

import "fmt"

type person struct {
	name     string
	nickname string
	age      uint8
}

type student struct {
	person
	university string
	course     string
}

func main() {
	s1 := student{
		person: person{
			name:     "Marcos",
			nickname: "Moreira",
			age:      27,
		},
		course:     "Computer Science",
		university: "MIT",
	}

	fmt.Println(s1)
	fmt.Println(s1.person)
	fmt.Println(s1.person.name)
}
