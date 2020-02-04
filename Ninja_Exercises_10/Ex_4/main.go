package main

import (
	"fmt"
)

// Exercise 4 Description:
// Starting with the original code, pull the values off the channel using a select statement

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

// We need to put the send-to channel part in a separate goroutine, which we do by wrapping the
// iterative part of the code in an anonymous goroutine. We also send a value to our channel q
// and close channel c when we finish sending values to it.
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

// For our receive function, we create a continuous for loop and utilize a select statement
// we have two cases:
//	1. The function pulls from channel c, then prints the value that it pulled
//	2. The function pulls from channel q, then the function returns, breaking out of our for loop
// In the latter case, the function will return to func main() and continue on with the code
func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
