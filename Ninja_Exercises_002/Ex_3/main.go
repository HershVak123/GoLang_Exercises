package main

import "fmt"

// Exercise 3 Description:
// Create TYPED and UNTYPED constants. Print the values of the constants.

// creating const with typed/untyped values
const (
	a     = 52
	b int = 53
)

func main() {

	// let's print our values
	fmt.Println(a, b)

	// let's print out our types as well just for curiosity's sake
	// They will both print out int when we run the code
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("b type: %T\n", b)
}
