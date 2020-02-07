package main

import "fmt"

// Exercise 4 Description:
// 1. Create a user defined struct with
//		- the identifier “person”
//		- the fields:
//			a. first
//			b. last
//			c. age
// 2. attach a method to type person with
//		- the identifier “speak”
//		- the method should have the person say their name and age
// 3. create a value of type person
// 4. call the method from the value of type person

// Let's first create our struct
type person struct {
	first string
	last  string
	age   int
}

// Now we will create a method that will take type person as a receiver.
// The way this works is, any value that is of type of the type in the receiver will have this method
// attached to it. That means any value of type person can use "Speak()" as a method.
func (p person) Speak() {
	fmt.Println("This is", p.first, p.last)
	fmt.Printf("They are %d years old\n", p.age)
}

func main() {

	// Create our person value
	p1 := person{
		first: "Hersh",
		last:  "Vakharia",
		age:   26,
	}

	// call the method that is attached to our value
	p1.Speak()
}
