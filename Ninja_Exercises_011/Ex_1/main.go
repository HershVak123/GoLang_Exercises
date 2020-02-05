package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Exercise 1 description:
// Starting with the original code, instead of using the blank identifier "_"
// for the error value, make sure the code is properly checking and handling the error

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

	// First, we will replace the blank identifier with "err" so that we can use it
	// Then we will use log.Fatalln(err) to stop the code immediately if there is an error
	// and print out the error.
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

}
