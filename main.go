package main

import "fmt"

func main() {

	// variable declare and initialize
	var message = "Hello World!"

	// types of string formatting
	// 1. s for simple
	// 2. q for qoute

	// show output
	fmt.Printf("Simple message = %s\n", message)
	fmt.Printf("Quote  message = %q\n", message)
}
