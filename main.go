package main

import "fmt"

func main() {

	// variable declaration
	var num1, num2 int

	// variable initialization by user input
	fmt.Printf("Enter the first number  : ")
	fmt.Scan(&num1)

	fmt.Printf("Enter the second number : ")
	fmt.Scan(&num2)

	// arithmetic operation using by + operator
	var sum = num1 + num2
	fmt.Printf("Summation of      : %v + %v = %v\n", num1, num2, sum)

	// arithmetic operation using by - operator
	var sub = num1 - num2
	fmt.Printf("Substruction of   : %v - %v = %v\n", num1, num2, sub)

	// arithmetic operation using by * operator
	var mul = num1 * num2
	fmt.Printf("Multiplication of : %v * %v = %v\n", num1, num2, mul)

	// arithmetic operation using by / operator
	var div = float32(num1) / float32(num2)
	fmt.Printf("Dividation of     : %v / %v : %v\n", num1, num2, div)

	// arithmetic operation using by % operator
	var mod = num1 % num2
	fmt.Printf("Modulation of     : %v / %v : %v\n", num1, num2, mod)
}
