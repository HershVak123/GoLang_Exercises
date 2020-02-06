package main

import "fmt"

// Exercise 2 Description:
// 1. Using a COMPOSITE LITERAL:
//		- create a SLICE of TYPE int
//		- assign 10 VALUES
// 2. Range over the slice and print the values out.
// 3. Using format printing
//		- print out the TYPE of the slice

func main() {

	// As opposed to arrays, where we must specify size, we create a slice by leaving the "size" parameter blank
	// in the brackets.
	// The advantage of using slices is that they can have values added and removed from them.

	// Here is our slice using a composite literal
	ss := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Ranging over the slice. If we use just "i", then we will only have the index available to us
	// So we use the blank identifier to tell our code that we don't want to do anything with the index,
	// but we do want to do something with the values
	for _, v := range sl {
		fmt.Println(v)
	}

	// Printing the type using fmt.Printf and %T
	// We will see that our slice is of type []int (slice of int)
	fmt.Printf("The type of our slice is: %T\n", s)
}
