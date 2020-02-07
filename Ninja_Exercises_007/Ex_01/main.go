package main

import "fmt"

// Exercise 1 Description:
// 1. Create a value and assign it to a variable.
// 2. Print the address of that value.

func main() {

	// assign value to variable
	x := 12

	// To print the address of a value, put '&' before the value in the print statement
	// This will give the memory address of the value stored in x
	fmt.Println("The address of x is:", &x)
}
