package main

import "fmt"

// Exercise 10 Description:
// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable

func main() {

	// We will now create 2 variables that are different instances of foo.
	f := foo()
	g := foo()

	// Since these are 2 separate instantiations, the incrementation of "x" will only be within the scope of that
	// instantiation. Let's print some values to show that

	// f()
	fmt.Println("x value in scope of f:", f())
	fmt.Println("x value in scope of f:", f())
	fmt.Println("x value in scope of f:", f())
	fmt.Println("x value in scope of f:", f())
	fmt.Println("x value in scope of f:", f())

	// g()
	fmt.Println("x value in scope of g:", g())
	fmt.Println("x value in scope of g:", g())
	fmt.Println("x value in scope of g:", g())
	fmt.Println("x value in scope of g:", g())
}

// What we will show in this exercise is that an variable can be defined in function-scope, and another function can
// manipulate the value of that variable within the scope of another function, but those manipulations will only exist
// in the scope of that function.

// We will define a variable in this function, but the return value will be a function that executes code
// to increment the value of that variable
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
