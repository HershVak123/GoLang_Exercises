package main

import "fmt"

// Exercise 3 Description:
// 1. Create a new type: vehicle.
//		- The underlying type is a struct.
//		- The fields:
//			a. doors
//			b. color
// 2. Create two new types: truck & sedan.
//		- The underlying type of each of these new types is a struct.
//		- Embed the “vehicle” type in both truck & sedan.
//		- Give truck the field “fourWheel” which will be set to bool.
//		- Give sedan the field “luxury” which will be set to bool. solution
// 3. Using the vehicle, truck, and sedan structs:
//		- using a composite literal, create a value of type truck and assign values to the fields;
//		- using a composite literal, create a value of type sedan and assign values to the fields.
// 4. Print out each of these values.
// 5. Print out a single field from each of these values.

// Let's first create our struct
type vehicle struct {
	doors int
	color string
}

// Now we create two new types with underlying type struct that "inherit" from vehicle
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	// Let's create our composite literals here
	// truck
	tr1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	// sedan
	se1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: true,
	}

	// Now we print out both of our structs
	fmt.Println(tr1)
	fmt.Println(se1)

	// Here we will print out one field from each struct
	// Since the field names from our inner struct don't conflict with the ones from our outer struct
	// we can call the fields from the inner struct directly with <struct>.<field>
	// as opposed to <outer>.<inner>.<field>
	fmt.Println(tr1.color)
	fmt.Println(se1.doors)

}
