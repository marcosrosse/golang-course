package main

// Interfaces

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

func writeArea(f form) {
	fmt.Printf("The area of the form is %0.2f\n", f.area())
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
