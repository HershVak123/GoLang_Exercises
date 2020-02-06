package main

import "fmt"

// Exercise 6 Description:
// Create a slice to store the names of all of the states in the United States of America.
// What is the length of your slice? What is the capacity?
// Print out all of the values, along with their index position in the slice, without using the range clause.

func main() {

	s := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`,
		` West Virginia`, ` Wisconsin`, ` Wyoming`}

	// to get the length of a slice, we use len(), to get the capacity we use cap()
	// Since slices are dynamically created arrays depending on the length of input we give, the length and capacity
	// will be the same
	// One situation it would be different is if we used a "make" command and specify the length and capacity
	// of the slice we are making (you always want cap > len)
	fmt.Println("The length our our slice is:", len(s))
	fmt.Println("The capacity of our slice is:", cap(s))

	// Iterate over the slice without range
	// we will implement the condition that i := 0, i < 50 to get our index
	// and we will use s[i] to print the value for each iteration

	for i := 0; i < 50; i++ {
		fmt.Printf("Index: %d\t State: %s\n", i, s[i])
	}
}
