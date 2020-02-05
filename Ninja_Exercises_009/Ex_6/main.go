package main

import (

	"fmt"
	"runtime"
)

// Exercise 6 Description:
// Create a program that prints out your OS and ARCH. Use the following commands to run it
//	- go run
//	- go build
//	- go install

func main() {

	fmt.Println("System OS:", runtime.GOOS)
	fmt.Println("System Architecture:", runtime.GOARCH)
}

// if we use go run on this file, it will execute the code from the file
// using go build will create a .exe file that runs the code
// go install will install this file to the bin directory, which is in our $PATH.
// As a result, we can just call the file name and it will execute