package main

import "fmt"

func main() {

	// Variable naming rules
	// 1. Letters, digits, _
	// 2. Keywords are not allowed as variable
	// 3. Variable can not have space
	// 4. Variable name can not start with digit

	// Variable declaration
	var firstName, lastName string
	var country string
	var age int
	var gpa float32

	// Variable initialization
	firstName = "Md."
	lastName = "Mohi-Uddin"
	country = "Bangladesh"
	age = 23
	gpa = 4.89

	// print output
	fmt.Println("My name is ", firstName, " ", lastName, " .")
	fmt.Println("I'm currently ", age, "years old")
	fmt.Println("I'm from ", country)
	fmt.Println("I have got GPA ", gpa, "in SSC")
}
