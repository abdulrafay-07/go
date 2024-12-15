package main

import (
	"fmt"
)

func getFactorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * getFactorial(number-1)
}

func main() {
	var factorial int

	fmt.Print("Enter a non-negative integer: ")
	_, err := fmt.Scan(&factorial)
	if err != nil || factorial < 0 {
		fmt.Println("Please enter a valid non-negative integer.")
		return
	}

	result := getFactorial(factorial)
	fmt.Printf("The factorial of %d is %d\n", factorial, result)
}
