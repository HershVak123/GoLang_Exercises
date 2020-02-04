package main

import (
	"fmt"
)

// Exercise 6 Description:
// write a program that
//	- puts 100 numbers to a channel
//	- pulls the numbers off the channel and print them

func main() {
	// First we initialize our channel to be general/bidirectional
	c := make(chan int)

	// We create a separate goroutine to run that will iteratively send values to our channel
	// Once the loop is over, the channel will close
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// This code here will pull/receive the values from our channel and print them
	for v := range c {
		fmt.Println(v)
	}
}

// When we run this code with go run -race main_test.go, we find that there are no race conditions
