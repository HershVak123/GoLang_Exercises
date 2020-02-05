package main

import "fmt"

// Exercise 9 Description:
// Create a program that uses a switch statement with the switch expression
// specified as a variable of TYPE string with the IDENTIFIER “favSport”.

func main() {

	// First we will create our variable with identifier "favSport"

	favSport := "Soccer"

	// next we will create a switch statement and use "favSport" as our switch expression
	switch favSport {

	// Notice for our cases that it's "case <possible value>:"
	// Since we assigned "Soccer", the first case will print.
	// If we change to another value, we will either get another case if the value is one of the listed cases
	// Or we will get the default case.
	case "Soccer":
		fmt.Println("Did you watch the Champion's League?")
	case "Basketball":
		fmt.Println("Kawhi Leonard is unstoppable right now!")
	case "Baseball":
		fmt.Println("I too enjoy a class-3 sedative...")
	case "Tennis":
		fmt.Println("Roger Federer is GOAT.")
	default:
		fmt.Println("What even is that sport?")
	}
}
