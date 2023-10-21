package main

import "fmt"

func main() {

	// Types of conditional statements => if, if-else, else-if, switch

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

	// switch statement
	switch number {
	case 0:
		fmt.Printf("The number is zero.\n")
	case 1:
		fmt.Printf("The number is one.\n")
	case 2:
		fmt.Printf("The number is two.\n")
	case 3:
		fmt.Printf("The number is three.\n")
	case 4:
		fmt.Printf("The number is four.\n")
	case 5:
		fmt.Printf("The number is five.\n")
	case 6:
		fmt.Printf("The number is six.\n")
	case 7:
		fmt.Printf("The number is seven.\n")
	case 8:
		fmt.Printf("The number is eight.\n")
	case 9:
		fmt.Printf("The number is nine.\n")
	case 10, 20, 30, 40, 50, 60, 70, 80, 90:
		fmt.Printf("The number is getter than 9 and less than 100.\n")
	case 100, 200, 300, 400, 500, 600, 700, 800, 900:
		fmt.Printf("The number is getter than 99 and less than 1000.\n")
	default:
		fmt.Printf("The number is %v\n", number)
	}

}
