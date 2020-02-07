package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Exercise 3 Description
// Starting with the original code, encode to JSON the []user sending the results to Stdout.
// Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	// The hint in the exercise description gives away 90% of the solution
	// The only thing we need to change is what goes in the Encode function
	// The answer to that is our []user variable
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

	// Reminder, since the output is being output from the code directly to whatever the output console is,
	// we can't really save the value to a variable and work with it.
	// If we run main.go we will see our output
}
