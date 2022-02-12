package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving the user: %s\n", u.name)
}

func (u user) major() bool {
	return u.age >= 18
}

func (u *user) anniversary() {
	u.age++
}

func main() {
	u1 := user{"Marcos", 27}
	fmt.Println(u1)
	u1.save()

	u2 := user{"Potato", 27}
	fmt.Println(u2)
	u2.anniversary()
	fmt.Println(u2.age)

}
