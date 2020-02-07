package main

import "fmt"

// Exercise 4 Description:
// Create and use an anonymous struct.

func main() {

	// to make use of an anonymous struct, we need to assign it to a variable
	// And then immediately invoke it with values

	c1 := struct {
		color  string
		doors  int
		luxury bool
	}{
		color:  "Grey",
		doors:  4,
		luxury: true,
	}

	// Let's print the fields from our struct
	fmt.Println(c1.color)
	fmt.Println(c1.doors)
	fmt.Println(c1.luxury)

	// Let's also print the type for good measure. Go is all about type after all
	fmt.Printf("type: %T\n", c1)
}
