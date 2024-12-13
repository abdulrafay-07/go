package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)

	// Prompt user for operation
	fmt.Print("Enter the operation you want to perform (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator) // Trim newline

	// Prompt user for numbers
	fmt.Print("Enter the first and the second number: ")
	var num1, num2 int
	_, err := fmt.Scanf("%d %d\n", &num1, &num2)
	if err != nil {
		fmt.Println("Error reading numbers:", err)
		return
	}

	switch operator {
	case "+":
		result := Add(num1, num2)
		fmt.Printf("%d + %d = %d", num1, num2, result)

	case "-":
		result := Subtract(num1, num2)
		fmt.Printf("%d - %d = %d", num1, num2, result)

	case "*":
		result, err := Multiply(num1, num2)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%d * %d = %d", num1, num2, result)

	case "/":
		quotient, remainder, err := Divide(num1, num2)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%d / %d = %d (remainder: %d)\n", num1, num2, quotient, remainder)

	default:
		fmt.Println("Invalid operator. Please use one of +, -, *, /.")
	}
}
