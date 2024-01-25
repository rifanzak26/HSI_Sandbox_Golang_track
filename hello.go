package main

import (
	"errors"
	"fmt"
)

func main() {
	// Create an error with a formatted message using fmt.Errorf
	err1 := fmt.Errorf("The number is not valid: %d", 1000)

	// Create an error with a static message using errors.New
	err2 := errors.New("The number is not valid")

	// Print the errors
	fmt.Println(err1) // Output: The number is not valid: 1000
	fmt.Println(err2) // Output: The number is not valid
}
