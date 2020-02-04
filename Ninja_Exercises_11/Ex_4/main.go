package main

import (
	"fmt"
	"log"
)

// Exercise 4 Description:
// Starting with the original code,  use the sqrt.Error struct as a value of type error.
//If you would like, use these numbers for your
//	- lat "50.2289 N"
//	- long "99.4656 W"

// Because our sqrtError type is received by the Error() func, it will utilize the error interface
// and as a result, the struct will be of type error
type sqrtError struct {
	lat  string
	long string
	err  error
}

// This function will take the lat/long and error values and return them in a formatted string
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

// because our sqrt function can return type error, that means it can return our sqrtError struct
// In the face that f < 0 (our condition for an error to occur), we define our error as a fmt.Errorf string,
// And then we call our sqrtError struct and provide the lat, long, and err values.
// And of course, in the event of no error, we will return 42 for our float64 value and nil for the error value
func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("custom error message: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", err}
	}
	return 42, nil
}
