package main

import "testing"

// We've imported our testing package, so we can now write our test.
// We create a TestCalculate function for the test, which will return an error if the test fails
// Otherwise it will pass the test
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2+2 to equal 4")
	}
}

// to actually run the test, we will go to our terminal, cd into this directory
// then use the "go test" command, and it will run the test and output:
// 	- how long the test took
// 	- the results of whatever tests needed to be ran