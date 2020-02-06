package main

import "fmt"

// Exercise 3 Description:
// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
//		[42 43 44 45 46]
//		[47 48 49 50 51]
//		[44 45 46 47 48]
//		[43 44 45 46 47]

func main() {

	// Since in the previous exercise, we used different values, we will create a new slice in this exercise
	// that contains all the relevant values above.

	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// slicing means that with our variable, we will choose a range from the index to work with
	// We do this with the notation <variable>[x:y]

	// For the first slice, we are capturing the first 5 values
	// Note: when slicing, the first value is included in the range, but the second value is not
	// That means we will need our slice to read [0:5], or just [:5]
	x := s[:5]
	fmt.Println(x)

	// The second slice looks like it ranges the last 5 values, so our slice will read [5:]
	y := s[5:]
	fmt.Println(y)

	// The third slice is from the 3rd to 7th value so our slice is [3:8]
	z := s[3:8]
	fmt.Println(z)

	// The last slice is from the 2nd to 6th value so our slice is [2:7]
	w := s[2:7]
	fmt.Println(w)
}
