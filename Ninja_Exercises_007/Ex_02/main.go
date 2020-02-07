package main

import "fmt"

// Exercise 2 Description
// 1. create a person struct
// 2. create a func called “changeMe” with a *person as a parameter
//		- change the value stored at the *person address
// 3. create a value of type person
//		- print out the value
// 4. in func main
//		- call “changeMe”
// 5. in func main
//		- print out the value

// First, let's create our struct
type person struct {
	name string
}

// Now we create our "changeMe" function
func changeMe(p *person) {
	// So the normally, we would have to dereference our p value before we call any field from it
	// ie. (*p).name,
	// but GoLang is nice, and doing just p.name is the shorthand for the above
	p.name = "Default Name"
}

func main() {

	// Create a value of type person
	p := person{name: "Hersh Vakharia"}

	// To show that we are indeed changing the value, let's print the value of p prior to calling changeMe
	fmt.Println(p)

	// Now, let's call changeMe
	// Remember, we are passing a pointer to a person as our parameter, so we need to add a '&' before the param.
	changeMe(&p)

	// Now we print the changed value
	fmt.Println(p)
}
