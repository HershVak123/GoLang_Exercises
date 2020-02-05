package main

import "fmt"

// Exercise 2 Description:
// Print every rune code point of the uppercase alphabet three times. Your output should look like this:
//	65
//		U+0041 'A'
//		U+0041 'A'
//		U+0041 'A'
//	66
//		U+0042 'B'
//		U+0042 'B'
//		U+0042 'B'

func main() {

	// For this solution we will be using a formatted string to print the unicode characters
	// The upper case characters are unicode values 65-90
	// Also since we are printing each character 3 times, we need to utilize a nested for loop in our
	// main for loop

	for i := 65; i <= 90; i++ {
		// This will print the non-tabbed unicode number
		fmt.Println(i)

		// This loop will print out the formatted string 3 times
		// The string-format symbol for unicode is %#U
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
