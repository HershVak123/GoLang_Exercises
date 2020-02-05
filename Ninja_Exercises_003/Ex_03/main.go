package main

import "fmt"

// Exercise 3 Description:
// Create a for loop using this syntax
//		- for condition { }
// Have it print out the years you have been alive.

func main() {

	// This requires a simple for loop
	// Our initial value will be the year 1993, the condition will be value < 2020
	// and we will increment value++

	for i := 1993; i < 2020; i++ {
		fmt.Println("I was alive in year", i)
	}
	// We intentionally made the condition < 2020 because I am currently alive in the year 2020 and will
	// make a print statement to indicate that
	fmt.Println("I am alive in year 2020")
}
