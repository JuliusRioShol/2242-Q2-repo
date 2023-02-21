// Interfaces
package main

import (
	"fmt"
	"math"
)

// Shape interface:
type Shape interface {
	area() float64
	circumf() float64
}

// 2 structs:
type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

// square method:
func (s Square) area() float64 {
	return s.length * s.length
}
func (s Square) circumf() float64 {
	return s.length * 4
}

// circle methods:
func (c Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// Make some general functions like these ones:
func printShapeInfo(s Shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())

}
func main() {
	square := Square{2}
	circle := Circle{5}
	printShapeInfo(square)
	printShapeInfo(circle)

}

/*
Interfaces:

Notes:
.An interface basically groups types together based on their methods
.Only provides method signatures


The use:
.Allow us to create some general functions


syntax:
type <nameOfInterface> interface {
	//method signatures.
}
*/
