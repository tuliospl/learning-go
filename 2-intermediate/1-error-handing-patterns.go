package main

import (
	"errors"
	"fmt"
)

// Error basic example
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Validating input example
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age: %d (needs to >= 0)", age)
	}
	if age > 150 {
		return fmt.Errorf("invalid age: %d (very high)", age)
	}
	return nil
}

func main() {
	// Example 1: Basic error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Validating input
	err = validateAge(-5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	err = validateAge(151)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}
}
