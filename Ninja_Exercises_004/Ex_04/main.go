package main

import "fmt"

// Exercise 4 Description
// 1. Follow these steps:
//		- start with this slice
//			x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//		- append to that slice this value
//			- 52
//		- print out the slice
//		- in ONE STATEMENT append to that slice these values
//			- 53
//			- 54
//			- 55
//		- print out the slice
//		- append to the slice this slice
//			- y := []int{56, 57, 58, 59, 60}
//		- print out the slice

func main() {

	// Start with our original slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// append to the slice. The notation is var = append(var, <new value(s)>). Then print.
	x = append(x, 52)
	fmt.Println(x)

	// Append multiple values at once. Just separate them with commas after our var in append(). Then print.
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	// create new slice and append to original
	// in order to append a second slice to a first, we need to "unfurl" the values of the second using "..." notation.
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
