package main

import "fmt"

// Exercise 1 Description:
// Write a program that prints a number in decimal, binary, and hex

func main() {
	x := 12345

	// We can represent the same int value in all of these ways with some string formatting
	// Since the original value is int, we can use %v to display the original value, or %d for base 10 (decimal)
	fmt.Printf("x in decimal: %d\n", x)

	// binary's string format is %b
	fmt.Printf("x in binary: %b\n", x)

	// hex's string format is $x (we can use %X to show the letters in capital)
	fmt.Printf("x in hex: %x\n", x)
}
