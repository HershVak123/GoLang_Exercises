package main

import "fmt"

// Exercise 2 Description:
// Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
// Access each value in the map. Print out the values, ranging over the slice.

// We'll copy over our code for defining our type person and the instantiation of our 2 people
type person struct {
	first   string
	last    string
	flavors []string
}

func main() {

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

	// Now we'll store the values in a map
	// Since we are using "last" name for the key and flavors for the value, we can just call those directly
	// From our structs using <var>.last and <var>.flavors
	m := map[string][]string{
		p1.last: p1.flavors,
		p2.last: p2.flavors,
	}

	// Now we iterate over the map to print out the keys and values
	for k, v := range m {
		fmt.Printf("Favorite flavors for last name %s:\n", k)

		for j, v2 := range v {
			fmt.Printf("\t%d. %s\n", j+1, v2)
		}
		fmt.Println()
	}
}
