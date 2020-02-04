package main

import (
	"fmt"
)

// Exercise 3 Description:
// Starting with the original code, pull the values off the channel using a for loop

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// The way that gen() was originally defined, it allowed for a deadlock to occur
// So what we needed to do was create a separate goroutine that executes within the function
// that sends values to our channel iteratively and closes the channel when we finish looping.
func gen() <-chan int {
	c := make(chan int)

	go func(){
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

// The receive function is a simple definition in which we take a value of a receive-only channel
// (Which was output from the gen function)
// and then we iterate over the range of the channel and print all the value out
func receive(c <-chan int) {
	for v := range c{
		fmt.Println(v)
	}
}

// This solution utilizes a the baton-pass methodology in which values will be sent to the channel in a
// separate goroutine, which blocks the routine, but allows the receive function in func main() to
// receive the values from the channel to print out