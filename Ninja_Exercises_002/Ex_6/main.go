package main

import "fmt"

// Exercise 6 Description:
// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

// We will use a const to define our iota.
// We start with the current year (2020) and iota for the next 4 years

const (

	// A cool bit of notation, if we use iota notation to define our first variable
	// then just list the subsequent variables, it is the same as saying
	// "variable = value + iota" for each subsequent variable
	a = 2020 + iota
	b
	c
	d
	e
)

func main() {

	// We now print the values
	// We will see a start with 2020 since iota will begin at 0
	// each subsequent iota will increment by 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
