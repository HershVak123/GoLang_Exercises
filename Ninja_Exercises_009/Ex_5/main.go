package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)
// Exercise 4 description:
// Fix the race condition caused in Ex 3 using Atomic

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	// Unlike Mutex, we don't need ot initialize a variable of Atomic to call its methods
	// We do, however, need to change our variable specifically to int64 because that is
	// a data type supported by Atomic. As such, we will use the other kind of
	// variable declaration to specify "counter" to be of type int64
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// Here, we will use Atomic's AddInt64 method to increment our variable. As you can see,
			// the method will take 2 arguments here: a pointer to our variable,
			// and the amount we wish to increment by.
			// This means that Atomic will update the value at the memory address directly.
			// To access our value, we will use LoadInt64, which only takes the pointer to our variable as an argument
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:\t",atomic.LoadInt64(&counter))

			wg.Done()
		}()

		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

// If we were to execute this code using go run -race main.go, we will find that there are no race conditions