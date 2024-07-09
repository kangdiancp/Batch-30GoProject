package math

import (
	"errors"
	"fmt"
)

func Multiply(a, b int) int {
	return a * b
}

// 1.handling errors with errors.New
func MultiplyOdd(n int) (int, error) {
	if n%2 == 0 {
		return 0, errors.New("n is not odd number")
	}
	return n * 2, nil
}

// 2.handling errors with fmt.ErrorF
func MultiplyEven(n int) (int, error) {
	if n%2 != 0 {
		return 0, fmt.Errorf("%d is not even number", n)
	}
	return n * 2, nil
}
