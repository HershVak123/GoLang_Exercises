package main

import (
	"encoding/json"
	"fmt"
)

// Exercise 1 Description:
// Starting with the original code, marshal the []user to JSON. There is a little bit of a curve ball in this
// hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

// First, we will create a struct
// For some reason, JSON doesn't like lower-case field names in the definition, so we either
// need to capitalize them, or use field tags (The former is much easier for our purposes)
type person struct {
	Name string
	Age  int
}

func main() {
	// Here we create 3 values of type person
	u1 := person{
		Name: "James",
		Age:  32,
	}

	u2 := person{
		Name: "Moneypenny",
		Age:  27,
	}

	u3 := person{
		Name: "M",
		Age:  54,
	}

	// Now we create a slice of person to store our person values
	people := []person{u1, u2, u3}

	fmt.Println(people)

	// your code goes here
	// the function we use to Marshal something to JSON is json.Marshal
	// we pass our []person into it as an argument
	// json.Marshal returns both the byte-string value of our []person as well as an error value
	// Good practice dictates we check this error value and take action if it is not nil
	// otherwise we just proceed normally with our program
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	// Since bs is a byte string, the printed value won't make much sense to us
	// so we will convert bs to string before printing it
	fmt.Println(string(bs))
}
