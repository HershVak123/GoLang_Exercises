package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples_es/008-ninja-level-twelve/01/dog"
)

type canine struct{
	name string
	age int
}

func main() {

	spot := canine{
		name: "Spot",
		age: dog.Years(7),
	}

	fmt.Printf("This is my dog %s! He is %d years old!\n", spot.name, spot.age)
}