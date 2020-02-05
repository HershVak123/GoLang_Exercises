package main

import "fmt"

func main() {

	// a raw string literal is wrapped in the backwards quote apostrophe
	// String literals will capture newlines and tabs and print them out.
	// We can also use double quotes in them too. How convenient!
	a := `This is a raw string literal.
			Every newline and tab will be accounted for
						in this.
		

			"We can even use double quotes in here"`

	// Now we print what we have
	fmt.Println(a)
}
