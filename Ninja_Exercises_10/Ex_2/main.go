package main

import (
	"fmt"
)

// Exercise 2 description:
// Make the following code work

func main() {
	// In the original code, the channel definition made the channel send-only
	// As a result, we are able to send values to it, but we can't receive values from it to print out
	// To fix this, we change the channel definition so that the channel is bidirectional
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
