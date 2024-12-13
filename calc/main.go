package main

import (
	"errors"
	"fmt"
)

func Add(n1, n2 int) int {
	return n1 + n2
}

func Subtract(n1, n2 int) int {
	return n1 - n2
}

func Multiply(n1, n2 int) (int, error) {
	if n1 > 1_000_000 && n2 > 1_000_000 {
		return 0, errors.New("number cannot be greater than 1000000")
	}

	return n1 * n2, nil
}

func Divide(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator cannot be equal to zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	fmt.Println("Testing arithmetic functions...")

	// Test Add
	fmt.Println("\n--- Add ---")
	fmt.Println("5 + 3 =", Add(5, 3))       // Normal case
	fmt.Println("-10 + 15 =", Add(-10, 15)) // Negative and positive
	fmt.Println("0 + 0 =", Add(0, 0))       // Zero case

	// Test Subtract
	fmt.Println("\n--- Subtract ---")
	fmt.Println("10 - 5 =", Subtract(10, 5))       // Normal case
	fmt.Println("-5 - (-10) =", Subtract(-5, -10)) // Double negatives
	fmt.Println("0 - 0 =", Subtract(0, 0))         // Zero case

	// Test Multiply
	fmt.Println("\n--- Multiply ---")
	result, err := Multiply(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 * 5 =", result) // Normal case
	}

	result, err = Multiply(1_000_001, 1_000_002)
	if err != nil {
		fmt.Println("Error:", err) // Overflow-like case
	} else {
		fmt.Println("1_000_001 * 1_000_002 =", result)
	}

	result, err = Multiply(0, 100)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("0 * 100 =", result) // Zero case
	}

	// Test Divide
	fmt.Println("\n--- Divide ---")
	quotient, remainder, err := Divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 3 =", quotient, "remainder =", remainder) // Normal case
	}

	quotient, remainder, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Divide by zero
	} else {
		fmt.Println("10 / 0 =", quotient, "remainder =", remainder)
	}

	quotient, remainder, err = Divide(0, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("0 / 5 =", quotient, "remainder =", remainder) // Zero numerator
	}

	quotient, remainder, err = Divide(-10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("-10 / 3 =", quotient, "remainder =", remainder) // Negative case
	}
}
