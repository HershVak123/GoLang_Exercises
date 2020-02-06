package main

import "fmt"

// Exercise 8 Description:
// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string
// which stores their favorite things. Store three records in your map. Print out all of the values,
// along with their index position in the slice.
//
//		- `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//		- `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//		- `no_dr`, `Being evil`, `Ice cream`, `Sunsets`

func main() {

	// First we create our map
	// To create a map we do: map[<key type>]<value type>{key: value, key: value, etc.}
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// Let's print our map to see what it looks like
	// We'll pritn an empty line after as well to not mess with the formatting of the next part of the exercise
	fmt.Println(m)
	fmt.Println("")

	// Now we have to iterate over the map to print out our key/value pairs as well as the
	// individual elements of the values

	for k, v := range m {
		fmt.Println("Favorite things for:", k)

		// Now we iterate over v, our []string
		for j, v2 := range v {
			// Use a formatted string to properly format the index value + 1 (to make a properly numbered list)
			// then we print our value at that index.
			fmt.Printf("\t%d. %v\n", j+1, v2)
		}
		fmt.Println("")
	}
}
