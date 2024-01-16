package main

import (
	"errors"
	"fmt"
)

// divide is a function that performs division of two numbers and returns the result and an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Create an error if the divisor is zero
		return 0, errors.New("cannot divide by zero")
	}
	// Return the result of the division and nil (no error)
	return a / b, nil
}

func main() {
	// Testing the divide function
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Testing the divide function again with values that don't cause an error
	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
