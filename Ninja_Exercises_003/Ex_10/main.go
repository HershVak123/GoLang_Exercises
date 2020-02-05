package main

import "fmt"

// Exercise 10 Description:
// Write down what these print:
//		- fmt.Println(true && true)
//		- fmt.Println(true && false)
//		- fmt.Println(true || true)
//		- fmt.Println(true || false)
//		- fmt.Println(!true)

func main() {

	// There are basic logic expressions, && means "and", || means "or", and ! means "not"

	// true and true is true
	fmt.Println(true && true)

	// true and false is false
	fmt.Println(true && false)

	// true or true is true
	fmt.Println(true || true)

	// true or false is true
	fmt.Println(true || false)

	// not true is false
	fmt.Println(!true)

}
