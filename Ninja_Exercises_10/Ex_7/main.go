package main

import (
	"fmt"
)

// Exercise 7 Description:
// write a program that
//	- launches 10 goroutines
//		- each goroutine adds 10 numbers to a channel
//	- pull the numbers off the channel and print them

func main() {

	// define our general channel
	c := make(chan int)

	// Since we are making 10 goroutines that each add 10 values to our channel, we will need to utilize
	// nested for loops.
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	// Since our channel will have repeated digits 0-9, we need to iterate over the length of the channel
	// which is 100 and then print out k, as well as the kth value in our channel that we receive.
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit code")
}
