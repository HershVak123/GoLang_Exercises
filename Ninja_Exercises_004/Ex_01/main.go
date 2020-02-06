package main

import "fmt"

//
// 1. Using a COMPOSITE LITERAL:
//		- create an ARRAY which holds 5 VALUES of TYPE int
//		- assign VALUES to each index position.
//		- Range over the array and print the values out.
// 2. Using format printing
//		- print out the TYPE of the array

func main() {
	// Here is where we define the array
	// The notation is [<size>]<type>{values}
	// If we leave the values blank, it will create the zero-array (an array of length <size> filled with 0's)
	// Due to the nature of arrays, we cannot add or remove values, we can change them, however
	// something something static programming

	x := [5]int{1, 2, 3, 4, 5}

	// We can use fmt.Printf to print the type of our array
	// We will see that the array is of type [5]int (which denotes array length 5 of int)
	fmt.Printf("The array's type is: %T\n", x)
}
