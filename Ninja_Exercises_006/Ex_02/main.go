package main

import (
	"fmt"
)

// Ex 2 Description:
// 1. create a func with the identifier foo that
//		- takes in a variadic parameter of type int
//		- pass in a value of type []int into your func (unfurl the []int)
//		- returns the sum of all values of type int passed in
// 2. create a func with the identifier bar that
//		- takes in a parameter of type []int
//		- returns the sum of all values of type int passed in

func main() {

	// We will create a []int then implement the two functions we created below
	xi := []int{2, 4, 6, 8, 10}

	//To send a slice through variadic param, we need ot "unfurl" it first
	f := foo(xi...)
	fmt.Println("foo sum:", f)

	//Since bar was defined to take []int, we don't need to unfurl the values
	b := bar(xi)
	fmt.Println("bar sum:", b)

}

// We will define our functions here.

// Let's start with the first function
// A variadic parameter is defined with "<param> ...<type>"
func foo(i ...int) int {
	sum := 0

	// We don't really need to use the index here, so we'll use the blank identifier
	for _, v := range i {
		sum += v
	}
	return sum
}

// Our second function will be almost identical
// The only difference is that i will be []int instead of a variadic parameter
func bar(i []int) int {
	sum := 0

	// We don't really need to use the index here, so we'll use the blank identifier
	for _, v := range i {
		sum += v
	}
	return sum
}
