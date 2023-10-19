package main

import "fmt"

func main() {

	// variable declaration
	var balance int

	// variable initialization
	fmt.Printf("Enter the balance : ")
	fmt.Scan(&balance)

	// integer to floating 32 bit type casting
	var floating32Balance = float32(balance)
	fmt.Println("Floating 32 balance is : ", floating32Balance)

	// integer to floating 64 bit type casting
	var floating64Balance = float64(balance)
	fmt.Println("Floating 64 balance is : ", floating64Balance)

}
