package main

import (
	"errors"
	"fmt"
)

// TODO: Implement these functions
// 1. Write a function 'add' that takes two integers and returns their sum
// 2. Write a function 'divide' that takes two integers and returns their quotient and an error if dividing by zero
// 3. Write a function 'calculatePower' that takes a base and exponent (integers) and returns the result

// Example solution for add (implement the rest yourself!):
func add(a, b int) int {
	return a + b
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a) / float64(b), nil
}

// Your implementation here
func calculatePower(base, exponent int) int {
	// TODO: Implement power calculation
	// Remember to handle negative exponents appropriately
	return 0
}

func main() {
	// Test your functions here
	fmt.Println("Testing add function:")
	fmt.Println("2 + 3 =", add(2, 3))

	fmt.Println("\nTesting divide function:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nTesting power function:")
	fmt.Println("2^3 should be 8, got:", calculatePower(2, 3))
} 