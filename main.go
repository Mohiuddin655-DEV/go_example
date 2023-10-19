package main

import "fmt"

func main() {

	// Types of logical operators => &&, ||, !

	// variable declaration
	var input int

	// variable initialization
	fmt.Print("Please enter the input number : ")
	fmt.Scan(&input)

	// use condition with logical operation
	var isAllConditionIsValid = input < 5 && input > 0
	var isAnyConditionIsValid = input < 5 || input > 0
	var isConditionUnvalid = !isAllConditionIsValid

	// show output
	fmt.Println("Logical operator (&&)  : ", isAllConditionIsValid)
	fmt.Println("Logical operator (||)  : ", isAnyConditionIsValid)
	fmt.Println("Logical operator (!)   : ", isConditionUnvalid)

}
