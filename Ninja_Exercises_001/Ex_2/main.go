package main

import "fmt"

// Exercise 2 Description:
// 1. Use var to DECLARE three VARIABLES. The variables should have package level scope.
// Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure
// the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
//		a. identifier “x” type int
//		b. identifier “y” type string
//		c. identifier “z” type bool
//2. in func main
//		a. print out the values for each identifier
//		b. The compiler assigned values to the variables. What are these values called?

// We will declare our variables outside of func main() so that they have package-level scope
// Since we are using the longer declaration method, we need to specify var and the data type in our declaration
var x int
var y string
var z bool

func main() {

	// Printing the values
	fmt.Println("x-value:", x)
	fmt.Println("y-value:", y)
	fmt.Println("z-value:", z)

	// Since we didn't explicitly state the values for our variables during their declaration
	// The program will automatically assign them the "zero-value". This will be different for each data type:
	//		- int/float will have a zero-value of 0
	//		- string will have a zero-value of the empty-string ""
	//		- bool will have a zero-value of false
}
