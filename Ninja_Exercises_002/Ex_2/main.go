package main

import "fmt"

// Exercise 2 Description:
// Using the following operators, write expressions and assign their values to variables:
//		g. ==
//		h. <=
//		i. >=
//		j. !=
//		k. <
//		l. >
//Now print each of the variables.

func main() {

	// Let's make the expressions and assign them to variables
	// all of these expressions will print out boolean values (true or false)
	g := (10 == 10)
	h := (7 <= 6)
	i := (5 >= 5)
	j := (40 != 40)
	k := (40 < 42)
	l := (40 > 42)

	fmt.Println("g value:", g)
	fmt.Println("h value:", h)
	fmt.Println("i value:", i)
	fmt.Println("j value:", j)
	fmt.Println("k value:", k)
	fmt.Println("l value:", l)
}
