package main

import (
	"fmt"
	"math"
)

// Exercise 5 Description:
// 1. create a type SQUARE
// 2. create a type CIRCLE
// 3. attach a method to each that calculates AREA and returns it
//		- circle area= π r 2
//		- square area = L * W
// 4. create a type SHAPE that defines an interface as anything that has the AREA method
// 5. create a func INFO which takes type shape and then prints the area
// 6. create a value of type square
// 7. create a value of type circle
// 8. use func info to print the area of square
// 9. use func info to print the area of circle

// First we will create our types. For now we will only give them values "side" for square and "radius" for circle
type square struct {
	side float64
}

type circle struct {
	radius float64
}

// Now we will create two separate area functions that take either circle or square type as receivers
func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Now we will create an interface for shape that will print out the info of whatever shape is input
type shape interface {
	area() float64
}

// now we will define our func info()
func info(s shape) {
	fmt.Println("The area of this shape is:", s.area())
}

func main() {

	// Let's define our cricle/square and then print the info
	c := circle{radius: 5.657}

	s := square{side: 7.49}

	// now we can call the info() to print out the area of each shape
	info(c)
	info(s)
}
