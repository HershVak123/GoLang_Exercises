package main

import (
	"fmt"
	"sync"
)

// Exercise 1 Description:
// In addition to the main goroutine, launch two additional goroutines
// 	- each additional goroutine should print something out
// 	- use waitgroups to make sure each goroutine finishes before your program exists

// First, we need to create a variable of type sync.WaitGroup
var wg sync.WaitGroup

func main() {

	// Since we have 2 separate go routines running, we need to add them both using wg.Add
	wg.Add(2)

	go foo()
	go okay()

	// Interesting note about the placement of wg.Wait(). If we put the Wait method here
	// the code will wait for both separate go routines to run before running bar()
	// Otherwise, if we placed the wait after bar(), bar would run first, then the separate go routines
	wg.Wait()
	bar()
}


func bar() {
	for i := 0; i < 10; i++{

		fmt.Println("bar:", i)
	}

}


// For both our functions that will be ran in separate go routines, we need to add wg.Done()
// to indicate to our code that what we need ran from the separate go routines has finished running

func foo() {
	for i := 0; i < 10; i++{

		fmt.Println("foo:", i)
	}
	wg.Done()
}

func okay() {
	for i := 0; i < 10; i++{

		fmt.Println("okay:", i)
	}
	wg.Done()
}