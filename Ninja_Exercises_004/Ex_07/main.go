package main

import "fmt"

// Exercise 7 Description:
// 1. Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//		- "James", "Bond", "Shaken, not stirred"
//		- "Miss", "Moneypenny", "Helloooooo, James."
//
// 2. Range over the records, then range over the data in each record.

func main() {

	// first we will create our slices of string for the two records above

	p1 := []string{"James", "Bond", "Shaken, not stirred"}
	p2 := []string{"Miss", "Moneypenny", "Hello, James"}

	// now we create our slice of slice of string
	s1 := [][]string{p1, p2}

	// Now we range over the entries in our [][]string and then range over the entries in those []strings
	// We iterate over both the index and value because we will use both in our loop
	for i, v := range s1 {
		// This will create a header of sort for each entry in our [][]string
		fmt.Printf("Record number %d:\n", i)

		// we won't use the index in this nested loop, but we will need the value v2
		// we iterate over range v (Which is our []string that has multiple entries itself)
		for _, v2 := range v {
			fmt.Println("\t", v2)

		}
		// We print an empty string after each [][]string entry for formatting
		fmt.Println("")
	}
}
