package main

import "fmt"

// Exercise 6 Description:
// Create a program that shows the “if statement” in action.

func main() {

	// We'll use a for loop for this exercise
	// We will iterate from 1 to 100 and only print values divisible by 3

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d is divisible by 3\n", i)
		}
	}
}
