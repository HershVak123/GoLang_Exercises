package main

import "fmt"

// Exercise 9 Description:
// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

func main() {

	// We'll copy our map from the previous example to save time
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// Now we need to add an entry to our map
	// To add an entry to a map, we use var[<key>] = <value>
	m[`McLeod_Todd`] = []string{`GoLang`, `Poetry`, `Learning`}

	// Now we iterate over our map and print the entries.
	// Since the code for this is the exact same as the previous exercise, we'll just copy it over

	for k, v := range m {
		fmt.Println("Favorite things for:", k)

		for j, v2 := range v {
			fmt.Printf("\t%d. %v\n", j+1, v2)
		}
		fmt.Println("")
	}
}
