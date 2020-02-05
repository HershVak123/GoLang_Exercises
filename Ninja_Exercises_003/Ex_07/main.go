package main

import "fmt"

// Exercise 7 Description:
// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

func main() {

	// For this example, we will build upon the code we wrote in Ex 6, but change a few things around
	// We will check for whether a number is divisible by 4, 3 or both.
	// If it's divisible by 4, we will print "red"
	// If it's divisible by 3, we will print "blue"
	// If divisible by both we will print "purple"
	// Otherwise we will just print the number

	for i := 1; i <= 100; i++ {
		// We begin with this condition first because if we begin with either of the individual
		// conditions, the code will never reach this point.
		// something something control flow.
		if (i%4 == 0) && (i%3 == 0) {
			fmt.Println("purple")
		} else if i%4 == 0 {
			fmt.Println("red")
		} else if i%3 == 0 {
			fmt.Println("blue")
		} else {
			fmt.Println(i)
		}
	}
}
