package main

import "fmt"

func main() {

	// variable declaration
	var balance float64

	// variable initialization
	fmt.Printf("Enter the balance : ")
	fmt.Scan(&balance)

	// floating number to integer type casting
	var balanceWithoutFloatingPoint = int(balance)
	fmt.Println("Floating number to integer balance : ", balanceWithoutFloatingPoint)

}
