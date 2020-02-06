package main

import "fmt"

// Exercise 5 Description:
// To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
// 	1. start with this slice
//		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	2. use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
//		[42, 43, 44, 48, 49, 50, 51]

func main() {

	// We start with our original slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Now we need to figure out the range of indices that give us the final slice above
	// We have 42 - 44, which is the first 3 values of the slice
	// and 48-51, which are the last 4 values

	// Remember when appending multiple slices together, we must unfurl each subsequent slice
	// after the first one
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}
