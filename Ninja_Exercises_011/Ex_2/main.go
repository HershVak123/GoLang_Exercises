package main

import (
	"encoding/json"
	"fmt"
)

// Exercise 2 Description:
// starting with the original code, Create a custom error message using "fmt.Errorf"

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Out here, we need to also run a check for whether error is nil or not
	// If err != nil, then we print the error and return from func main(),
	// thus exiting our code.
	// Conversely, we could use our options from the log package as well, depending on the situation.
	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
// To do this, we change the truen values to ([]byte, error)
// Then, we replace the blank identifier with "err", run a check to see if error is not nil
// if err != nil, we return an empty slice of byte and our fmt.Errorf error
// Otherwise we return the Marshal'd value and nil, since we have no error
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("JSON did not Marshall: %v\n", err)
	}

	return bs, nil
}
