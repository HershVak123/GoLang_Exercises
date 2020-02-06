package main

import "fmt"

// Exercise 10 Description:
// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

func main() {

	// We will again copy the same map over, but with the 4th entry added this time
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
		`McLeod_Todd`:     []string{`GoLang`, `Poetry`, `Learning`},
	}

	// Now we need to delete an entry from our map. Let's choose the `bond_james` entry.
	// To delete an entry, we do: delete(<map>, <key to delete>)
	delete(m, `bond_james`)

	// Now we iterate over our map to check if the entry was deleted
	// Again, the iteration code is the same as the previous two exercises, so we just copy it over
	for k, v := range m {
		fmt.Println("Favorite things for:", k)

		for j, v2 := range v {
			fmt.Printf("\t%d. %v\n", j+1, v2)
		}
		fmt.Println("")
	}
}
