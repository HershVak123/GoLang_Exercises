package main

import "fmt"

// Exercise 1 Description:
// Create your own type “person” which will have an underlying type
// of “struct” so that it can store the following data:
//		- first name
//		- last name
//		- favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

// First, let's create our person struct
// The syntax is "type <name> struct{<data>}
// for each entry in our struct, we need to specify the value name and data type
// for "first" and "last" both will be strings, but since ice cream flavors will have multiple values
// We will use type []string
type person struct {
	first   string
	last    string
	flavors []string
}

func main() {

	// Now we need to create two variables that have values of type person
	// When we instantiate a struct, we need to separate the value/identifier pairs with a comma
	p1 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"Lemon", "Lime", "Chocolate"},
	}

	p2 := person{
		first:   "Hersh",
		last:    "Vakharia",
		flavors: []string{"Chocolate", "Vanilla", "Cookie Dough"},
	}

	// We will need to iterate over the flavors value to print them out nicely
	// We will need to do this for each person
	fmt.Println("Favorite flavors for", p1.first, p1.last, ":")
	for i, v := range p1.flavors {
		fmt.Printf("\t%d. %s\n", i+1, v)
	}

	fmt.Println()

	// Printing values for p2
	fmt.Println("Favorite flavors for", p2.first, p2.last, ":")
	for i, v := range p2.flavors {
		fmt.Printf("\t%d. %s\n", i+1, v)
	}
}
