package main

import (
	"fmt"
	"sort"
)

// Exercise 5 Description
// 1. Starting with the original code, sort the []user by
//		- age
//		- last
// 2. Also sort each []string “Sayings” for each user
//		- print everything in a way that is pleasant

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// In Order to sort values in a struct, we need to create a custom sort method.
// We will accomplish this by creating a type which is a []user and attach methods to those []user that are
// the sorting rules for those []user

// Let's start by creating one to sort by age
type ByAge []user

// A function that returns the length of the []user
func (a ByAge) Len() int { return len(a) }

// A function that will swap items in our []user based on a condition
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// A function that returns a boolean whether the value at pos i is less than the value as pos j
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Now let's create one to sort by the Last field
// The functions are more or less the same, except here we will use a.Last rather than a.Age

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

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
	// Now we can use the sorting rules that we created above, we'll combine the "sort the sayings" step
	// of the exercise with sorting everything else just to be fancy and save some space

	// Sort ByAge
	sort.Sort(ByAge(users))
	for i, v := range users {
		fmt.Printf("Person %d: %s %s\n", i, v.First, v.Last)
		fmt.Printf("\t%s %s's age: %d\n", v.First, v.Last, v.Age)
		fmt.Printf("\t%s %s's sayings:\n", v.First, v.Last)
		sort.Strings(v.Sayings)
		for k, v2 := range v.Sayings {
			fmt.Printf("\t%d. %s\n", k+1, v2)
		}
		// Empty line for formatting
		fmt.Println()
	}

	fmt.Println()

	// Sort ByLast
	sort.Sort(ByLast(users))
	for i, v := range users {
		fmt.Printf("Person %d: %s %s\n", i, v.First, v.Last)
		fmt.Printf("\t%s %s's age: %d\n", v.First, v.Last, v.Age)
		fmt.Printf("\t%s %s's sayings:\n", v.First, v.Last)
		sort.Strings(v.Sayings)
		for k, v2 := range v.Sayings {
			fmt.Printf("\t\t%d. %s\n", k+1, v2)
		}
		// Empty line for formatting
		fmt.Println()
	}
}
