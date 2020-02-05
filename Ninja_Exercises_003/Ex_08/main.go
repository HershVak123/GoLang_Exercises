package main

import "fmt"

// Exercise 8 Description:
// Create a program that uses a switch statement with no switch expression specified.

func main() {

	// Since we have no switch expression, we will create a very basic switch statement.
	switch {

	// We will use true and false cases to show that the true case will always print
	// and that the false case will never print (unless we use fallthrough, but that's another discussion entirely)
	case true:
		fmt.Println("This statement will print.")

	case false:
		fmt.Println("This statement will never print.")
	}
}
