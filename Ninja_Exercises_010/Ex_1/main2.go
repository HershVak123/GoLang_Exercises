package main

import (
	"fmt"
)

// Exercise 1 Description:
// Get this code out of deadlock

// This code will implement buffer solution

func main() {
	// The only thing we need to do here is add a value to represent the number of values
	// we want to buffer in our channel regardless of whether or not they are
	// ready to be pulled
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
