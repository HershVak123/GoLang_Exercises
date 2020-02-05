package main

import(
	"fmt"
)

// Exercise 2 Description:
// This exercise will reinforce our understanding of method sets:
//	- create a type person struct
//	- attach a method speak to type person using a pointer receiver
//		- *person
//	- create a type human interface
//		- to implicitly implement the interface, a human must have the speak method
//	- create func “saySomething”
//		- have it take in a human as a parameter
//		- have it call the speak method
//	- show the following in your code
//		- you CAN pass a value of type *person into saySomething
//		- you CANNOT pass a value of type person into saySomething

type person struct{
	Name string
}

func (p *person) speak() {
	fmt.Println("Person is speaking:", p.Name)
}

type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}

func main() {

	p1 := person{Name:"Hersh Vakharia"}

	// Since our function speak has a receiver of type "pointer to person", our human interface attaches to this method
	// and since "saySomething" takes type human as an argument, using transitive property
	// we can deduce that saySomething will only accept pointer to person as an argument
	saySomething(&p1)

	// saySomething(p1)
	// We can uncomment the above line of code and run it, but we will get an error
	// because saySomething only accepts value of type pointer to person and not type person.

}