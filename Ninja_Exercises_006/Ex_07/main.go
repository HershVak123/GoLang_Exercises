package main

import "fmt"

// Exercise 7 Description:
// Assign a func to a variable, then call that func

func main() {

	// To assign an anonymous function to a variable, we just do
	// <var> := func(){<code>} without the execution () at the end.
	// Including the () at the end means that we would be storing the value of that function in the variable
	// rather than the function itself
	f := func() {
		fmt.Print("Anonymous function stored in a variable")
	}

	// now that our function is stored in a variable, we can execute it using f()
	f()
}
