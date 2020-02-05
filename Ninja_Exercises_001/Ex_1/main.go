package main

import "fmt"

// Exercise 1 Description:
// 1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 		a. 42
//		b. "James Bond"
//		c. true
// 2. Now print the values stored in those variables using
// 		a. a single print statement
// 		b. multiple print statement

func main() {

	// Since we're using the short declaration operator, we won't need to specify type, as Go will decide it for us.

	x := 42           // type int
	y := "James Bond" // type string
	z := true         // type bool

	// Printing in one print statement
	fmt.Println("Our three values:", x, y, z)

	// Printing over multiple statements
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
