package main

import "fmt"

// Exercise 5 Description:
// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

func main() {

	// Let's use a for loop for this
	// The loop will begin at i := 10, have condition i <= 100, and then increment i++

	for i := 10; i <= 100; i++ {

		// We'll store the remainder in a variable because we will use it in a formatted string
		rem := i % 4
		fmt.Printf("The remainder when %d is divided by 4 is: %d\n", i, rem)
	}
}
