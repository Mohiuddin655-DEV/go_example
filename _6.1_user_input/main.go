package main

import "fmt"

func main() {

	// variable declaration
	var firstName, lastName string
	var country string
	var age int
	var gpa float32

	// use two signs to get user input
	// 1. & sign used for single user input
	// 2. %v sign used for multiple user input

	// user input by scanf, scanln, scan (variable initialization)

	fmt.Print("Enter the full name : ")
	fmt.Scanf("%v %v", &firstName, &lastName)

	fmt.Print("Enter the country name : ")
	fmt.Scanln(&country)

	fmt.Print("Enter the age : ")
	fmt.Scanln(&age)

	fmt.Print("Enter the GPA : ")
	fmt.Scan(&gpa)

	// show output
	fmt.Printf("My name is %v %v.\n", firstName, lastName)
	fmt.Printf("I'm currently %v years old\n", age)
	fmt.Printf("I'm from %v\n", country)
	fmt.Printf("I have got GPA %v in SSC\n", gpa)
}
