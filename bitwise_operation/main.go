package main

import "fmt"

func main() {

	// Types of bitwise operators => &, |, ^, >>, <<

	var num1 = 10
	var num2 = 5

	fmt.Println("Bitwise operator (&)  : ", num1&num2)
	fmt.Println("Bitwise operator (|)  : ", num1|num2)
	fmt.Println("Bitwise operator (^)  : ", num1^num2)
	fmt.Println("Bitwise operator (>>) : ", num1>>num2)
	fmt.Println("Bitwise operator (<<) : ", num1<<num2)
}
