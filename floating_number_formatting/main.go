package main

import "fmt"

func main() {

	// variable declare and initialize
	var number = 3.1416

	// types of string formatting
	// 1. f for integer to float
	// 2. .2f for floating point first 2 digits

	// show output
	fmt.Printf("Current value for floating type = %f\n", number)
	fmt.Printf("Last two decimal value = %.2f\n", number)
}
