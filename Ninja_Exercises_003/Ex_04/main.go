package main

import "fmt"

// Exercise 4 Description:
// Create a for loop using this syntax
//		- for { }
// Have it print out the years you have been alive.

func main() {

	// Since we will be utilizing an infinite for loop, we need to
	// implement a condition for which the loop will break

	// We first need to initialize a variable with the starting year
	year := 1993

	for {
		// This is our break condition
		if year > 2020 {
			break
		}

		// The loop will continuously print the value of "year" then increment it by 1
		// until it hits the break condition.
		fmt.Println(year)
		year++
	}
}
