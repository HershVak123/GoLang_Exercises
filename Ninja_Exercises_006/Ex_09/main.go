package main

import "fmt"

// Ex 9 Description:
// A “callback” is when we pass a func into a func as an argument. For this exercise,
//		- pass a func into a func as an argument

func main() {
	// Our inner function will be an anonymous function that sums up the values of a []int
	// we will save the function to a variable
	g := func(yi []int) int {
		sum := 0

		// Iterate over the values of yi, ignoring the index
		for _, v := range yi {
			sum += v
		}
		return sum
	}

	// Create a []int for our input

	si := []int{2, 4, 6, 8, 10}

	// now we can see callback in action
	x := foo(g, si)

	fmt.Println("The results of our callback is the value:", x)
}

// We'll create our func that takes another func as an argument here
// To properly utilize callback, we need to describe what arguments the arg func takes and
// what type it returns

// Let's make our outer function take a function and slice of int as our two arguments
// and let's make our inner function take []int as an argument
// For the function argument, we have to specify the variable and type of value taken
// as well as the return value type.
// Since the inner function takes []int as an argument and returns type int, we will specify that
func foo(f func(yi []int) int, xi []int) int {
	// Here, the outer func will execute the inner func, and save the return value to another variable
	ret := f(xi)

	// We will do a simple manipulation with the value
	return ret * 2
}
