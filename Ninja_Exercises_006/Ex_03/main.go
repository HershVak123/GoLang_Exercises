package main

import "fmt"

// Exercise 3 Description:
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

func main() {

	// Here we will call foo() and bar(), but we will defer foo()

	defer foo()
	bar()

	// Since foo() is listed first, normally it would run first
	// However, since we placed "defer" before foo(), the program will wait until
	// Just before exiting to run it. As a result, foo() will run after bar()
	// We will see this when the program runs

}

// We will create two simple functions to show how defer works.
// These functions will just print which function they are from

func foo() {
	fmt.Println("Printing from foo")
}

func bar() {
	fmt.Println("Printing from bar")
}
