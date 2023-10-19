package main

import "fmt"

func main() {

	// variable declaration
	var number int

	fmt.Printf("Enter the number : ")
	fmt.Scan(&number)

	// types of number formatting
	// 1. b for binary
	// 2. o for octal
	// 3. x for hexa

	// show output
	fmt.Printf("Binary number = %b\n", number)
	fmt.Printf("Octal  number = %o\n", number)
	fmt.Printf("Hexa   number = %x\n", number)
}
