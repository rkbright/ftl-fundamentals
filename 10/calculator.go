// Package calculator module
package calculator

import "errors"

// Add returns sum of two ints
func Add(a, b int) int {
	return a + b
}

// Subtract two numbers
func Subtract(a, b int) int {
	return a - b
}

// Multiply two numbers
func Multiply(a, b int) int {
	return a * b
}

// Divide two numbers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}
