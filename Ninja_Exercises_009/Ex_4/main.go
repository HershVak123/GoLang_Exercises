package main

import (
	"fmt"
	"runtime"
	"sync"
)
// Exercise 4 description:
// Fix the race condition caused in Ex 3 using Mutex

// First, we need to create a variable of type sync.Mutex
var mu sync.Mutex

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// Here, we will surround the code that is affected by the Data Race with Mutex's Lock and Unlock methods
			// we will call mu.Lock() right before the relevant code executes and mu.Unlock() after
			// it is done executing.
			// This will lock the code so that only 1 Go Routine can access it at a time.
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()

			wg.Done()
		}()

		fmt.Println("GoRoutines:", runtime.NumGoroutine())
		fmt.Println("Counter:", counter)
	}
	wg.Wait()

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

// If we were to execute this code using go run -race main.go, we will find that there are no race conditions