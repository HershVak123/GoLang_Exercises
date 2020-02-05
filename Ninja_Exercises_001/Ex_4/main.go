package main

import "fmt"

// Exercise 4 Description:
//	1. Create your own type. Have the underlying type be an int.
//	2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
//	3. in func main
//		a. print out the value of the variable “x”
//		b. print out the type of the variable “x”
//		c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
//		d. print out the value of the variable “x”

// create type "numeral" with underlying type int
type numeral int

// define variable x of type numeral
var x numeral = 100

func main() {

	// printing the value and then type of x
	fmt.Println("x value:", x)
	fmt.Printf("X type: %T\n", x)

	// we can assign 42 to x because x's underlying type is int, and 42 is an int, therefore assignment can
	// work like this for underlying types
	x = 42

	fmt.Println("x value after assignment:", x)

	// Let's print the type again to show that x's type remains the same after assignment
	fmt.Printf("x type after assignment: %T\n", x)

}
