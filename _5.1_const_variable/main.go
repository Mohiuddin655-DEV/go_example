package main

import "fmt"

func main() {

	// const variable declaration naming rules
	// 1. must be use const keywork before
	// 2. use uppercase const variables but not mandatory

	// const variable
	const COUNTRY = "Bangladesh"
	const firstName string = "Md."
	const lastName string = "Mohi-Uddin"
	const age int = 23
	const gpa float32 = 4.89

	// show output
	fmt.Printf("My name is %v %v.\n", firstName, lastName)
	fmt.Printf("I'm currently %v years old\n", age)
	fmt.Printf("I'm from %v\n", COUNTRY)
	fmt.Printf("I have got GPA %v in SSC\n", gpa)
}
