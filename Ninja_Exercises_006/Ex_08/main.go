package main

import "fmt"

// Exercise 8 Description:
// 1. Create a func which returns a func
// 2. assign the returned func to a variable
// 3. call the returned func

func main() {

	// Since bar() is going to return a function, but not execute it, let's store that function in a variable
	f := bar()

	// At this point, we should see the print statement from bar(), but not from foo just yet
	// This is because foo has been made the value of a function. When we execute it (using '()')
	// the statement in foo will print
	// To show this, we will print a few blank lines before calling foo to show this behavior
	fmt.Println()
	fmt.Println()
	fmt.Println()
	f()
}

// Let's create the func that will be returned by the outer func
func foo() {
	fmt.Println("foo is now printing")
}

// Now we create the func that will return the inner func
func bar() func() {
	fmt.Print("We are about to return foo")
	return foo
}
