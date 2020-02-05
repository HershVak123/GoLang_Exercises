package main

import "fmt"

// Exercise 5 Description:
// 	1. Building on the code from the previous example
//	at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
//	The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// 	2. in func main
//		a. this should already be done
//			i. print out the value of the variable “x”
//			ii. print out the type of the variable “x”
//			iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//			iv. print out the value of the variable “x”
//		b. now do this
//			i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
//				1. then use the “=” operator to ASSIGN that value to “y”
//				2. print out the value stored in “y”
//				3. print out the type of “y”

// create type "numeral" with underlying type int
type numeral int

// define variable x of type numeral and var y as type int
var x numeral = 100
var y int

func main() {

	// This block of code is the same as Exercise 4
	fmt.Println("x value:", x)
	fmt.Printf("X type: %T\n", x)
	x = 55
	fmt.Println("x value after assignment:", x)

	// Here we will use conversion to assign our x value to y
	// We can't do a direct assignment because x and y are of different types (even if the subtypes are the same)
	y = int(x)

	// now we print the value to show that the conversion was successful
	// we also print the type for y to show that y is still of type int
	fmt.Println("y value:", y)
	fmt.Printf("y type: %T\n", y)

}
