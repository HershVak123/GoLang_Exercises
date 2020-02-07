package main

import (
	"encoding/json"
	"fmt"
)

// Exercise 2 Description
// Starting with the original code, unmarshal the JSON into a Go data structure

// Create a struct to store the Unamrshal'd values
type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var people []person

	// The way json.Unmarshal works is that the function takes the []byte value of what you want to Unmarshal
	// and a pointer to a data structure you want to store the data in (In this case, a []person)
	// This means that the function changes the value at the memory address of the []person directly,
	// so it won't return a value for us to store in a variable for the data structure.
	// The function still returns an error value, so we do need to check that for the sake of good practice.

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	// Now we need to iterate over the struct and print stuff out
	// Since our struct contains a field that is []string, we need to iterate over that too
	for i, v := range people {
		fmt.Println("Person", i)
		fmt.Println("\tPerson's name:", v.First, v.Last)
		fmt.Printf("\t%s %s's age: %d\n", v.First, v.Last, v.Age)
		fmt.Printf("\t%s %s's sayings:\n", v.First, v.Last)

		for k, v2 := range v.Sayings {
			fmt.Printf("\t\t%d. %s\n", k+1, v2)
		}
		// For formatting
		fmt.Println()
	}
}
