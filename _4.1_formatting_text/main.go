package main

import "fmt"

func main() {

	// variable declaration and initialization
	firstName := "Md."
	lastName := "Mohi-Uddin"
	country := "Bangladesh"
	age := 23
	gpa := 4.89

	// show output with formatting text
	fmt.Printf("My name is %v %v.\n", firstName, lastName)
	fmt.Printf("I'm currently %v years old\n", age)
	fmt.Printf("I'm from %v\n", country)
	fmt.Printf("I have got GPA %v in SSC\n", gpa)
}
