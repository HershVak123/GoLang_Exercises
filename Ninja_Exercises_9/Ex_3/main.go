package main

import (
	"fmt"
	"runtime"
	"sync"
)
// Exercise 3 description:
// Using goroutines, create an incrementer program
//	- have a variable to hold the incrementer value
//	- launch a bunch of goroutines
//	- each goroutine should
//		- read the incrementer value
//			- store it in a new variable
//		- yield the processor with runtime.Gosched()
//		- increment the new variable
//		- write the value in the new variable back to the incrementer variable
//	- use waitgroups to wait for all of your goroutines to finish
//	- the above will create a race condition.
//	- Prove that it is a race condition by using the -race flag

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

// If we were to execute this code with go run -race main_test.go
// we would see that there is at least one "Data Race"