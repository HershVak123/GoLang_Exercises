package main

import (
	"fmt"
)

// Exercise 3 Description:
//Create a struct “customErr” which implements the builtin.error interface.
//Create a func “foo” that has a value of type error as a parameter.
//Create a value of type “customErr” and pass it into “foo”.


// First step, create our customErr struct, and give it "info" which will be a string that
// describes our error
type customErr struct {
	info string
}

// We then define a function Error() that takes a receiver of type customError and will output a string containing
// our error message
// Since the builtin.error interface will assign type error to anything associated with an Error() method,
// Our customErr struct will be of type error (because polymorphism and whatnot)
func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	// Here, we will assign a variable to our struct of type customErr.
	c := customErr{"This is a custom error"}
	fmt.Printf("%T\n", c)

	// Because of the description above, we can pass our struct into function foo and will be accepted as
	// a parameter of type error. If we run our code, it will print everything out.
	foo(c)
}

// Here, we define foo with a parameter of type error. The function will print out the error with
// an additional string before it.
func foo(e error) {
	fmt.Println("foo ran -", e)
}
