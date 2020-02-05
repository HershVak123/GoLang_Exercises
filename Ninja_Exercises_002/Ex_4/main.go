package main

import "fmt"

// Exercise 4 Description:
// Write a program that
//	- assigns an int to a variable
//	- prints that int in decimal, binary, and hex
//	- shifts the bits of that int over 1 position to the left, and assigns that to a variable
//	- prints that variable in decimal, binary, and hex

func main() {

	// declaring our int variable and printing the value in the required formats
	a := 100
	fmt.Printf("a in decimal: %d\n", a)
	fmt.Printf("a in binary: %b\n", a)
	fmt.Printf("a in hex: %x\n", a)

	// shifting bits over to the left by 1 then printing the new value in the required formats
	b := a << 1
	fmt.Printf("b in decimal: %d\n", b)
	fmt.Printf("b in binary: %b\n", b)
	fmt.Printf("b in hex: %x\n", b)
}
