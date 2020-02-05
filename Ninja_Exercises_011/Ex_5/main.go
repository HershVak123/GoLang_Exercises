package main

import "fmt"

// Exercise 5 Description:
// We are going to learn about testing next.
// For this exercise, take a moment and see how much you can figure out about testing
// by reading the testing documentation & also Caleb Doxsey’s article on testing.
//See if you can get a basic example of testing working.

// For this exercise we will create a function that we wil ltest with in our test file.
// We will keep it simple and make a function that adds 2 to whatever value we input and returns the value
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

// Call the function in func main()
func main() {
	fmt.Println("Testing")
	result := Calculate(2)
	fmt.Println(result)
}

// To do the actual testing, we need to set up our test in our test file
// We created a file called main_test.go (<filename>_test.go seems to be the standard)