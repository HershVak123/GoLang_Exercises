package main

import (
	"fmt"
)

// Exercise 1 Description:
// Get this code out of deadlock

// This code will implement the baton-pass solution of creating a separate go routine for
// sending the value to our channel

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
