package main

import (
	"fmt"
)

// Exercise 2 Description:
// make the following code work

func main() {
	// As opposed to the other version of this code, which defined the channel as send-only
	// This version of the code defined the channel to be receive only, meaning that we can't send any values
	// into it.
	// The solution here is more or less the same here:
	// Change the channel definition to be bidirectional
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
