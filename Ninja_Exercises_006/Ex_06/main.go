package main

import "fmt"

// Exercise 6 Description
// Build and use an anonymous func

func main() {

	// To create an anonymous function, you just say func(){<code to execute>}()
	// be sure to use the () at the end to ensure the func runs, otherwise you won't be able to do anything with it
	// unless you assign it to a variable (which is going to be the next exercise)
	func() {
		fmt.Print("This is an anonymous function executing.")
	}()
}
