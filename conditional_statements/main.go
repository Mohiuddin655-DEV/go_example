package main

import "fmt"

func main() {

	// Types of conditional statements => if, if-else, else-if

	// Variable declaration
	var number int

	// Variable initialization
	fmt.Printf("Enter the number : ")
	fmt.Scan(&number)

	// if statement
	if number < 0 {
		fmt.Println("The number is less than 0.")
	}

	// if-else statement
	if number > 0 && number < 10 {
		fmt.Println("Hello World!")
	} else {
		fmt.Println("Number isn't valid!")
	}

	// else-if statement
	if number == 0 {
		fmt.Println("The number is 0.")
	} else if number > 0 && number < 10 {
		fmt.Println("The number is getter than 0 and less than 10!")
	} else {
		fmt.Printf("The entired number was %v\n", number)
	}

}
