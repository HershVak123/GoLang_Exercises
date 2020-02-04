package main

import (
	"fmt"
)

// Exercise 5 Description:
// Show the comma okay idiom for the following code

func main() {
	c := make(chan int)

	// Use a separate goroutine to send a value to the channel
	go func(){
		c <- 142
	}()

	// Now that the concurrent goroutine is running and has blocked, func main() is ready to receive
	// the value from our channel.
	// We use the comma okay idiom to check if the channel has values This first iteration we will see that the
	// value that we sent to the channel will be received and will print, and that the channel indeed had a value.
	v, ok := <-c
		fmt.Println(v, ok)

	close(c)

	// We closed the channel now, and since we already pulled the value that was sent to it
	// If we were to perform this check again, we will find that there are no values in the channel
	// Which will be denoted with a "0, false" for our v, ok values, respectively.
	v, ok = <-c
		fmt.Println(v, ok)
}
