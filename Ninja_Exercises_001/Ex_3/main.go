package main

import "fmt"

// Exercise 3 Description:
// Using the code from the previous exercise,
// 1. At the package level scope, assign the following values to the three variables
//		a. for x assign 42
//		b. for y assign “James Bond”
//		c. for z assign true
// 2. in func main
//		a. use fmt.Sprintf to print all of the VALUES to one single string.
//		   ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE
//		   with the IDENTIFIER “s”
//		b. print out the value stored by variable “s”

// Assigning variables at package-level scope
var x = 42
var y = "James Bond"
var z = true

func main() {

	// fmt.Sprintf will return a formatted string of whatever we input
	// Since we formatted the string to have newlines and %v to format values into, those
	// will show up in our nicely formatted string
	// s will be of type string, so the variable types of x y and z can't be extracted from it
	s := fmt.Sprintf("x value: %v\ny value: %v\nz value: %v\n", x, y, z)

	// Since we saved our formatted string to a variable, we can just print it using fmt.Println
	fmt.Print(s)
}
