package main

import (
	"fmt"
	"sort"
)

// Exercise 4 Description:
// Starting with the original code, sort the []int and []string for each person.

func main() {

	// Our []int and []string
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry",
		"fragmented", "moons", "across", "magenta"}

	// To show that the slices actually sort, we will print the value pre-sort and post-sort for each slice

	// []int
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	// []string
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
