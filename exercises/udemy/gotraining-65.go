/*
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/
package main

import (
	"fmt"
	"math"
)

// interface and function
type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("area of %T: %v\n", s, s.area())
}

// custom types
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// custom methods
func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := circle{
		radius: 2.15,
	}

	s1 := square{
		length: 4.22,
	}

	info(c1)
	info(s1)
}
