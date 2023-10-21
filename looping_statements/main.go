package main

import "fmt"

func main() {

	// Types of looping statements => for

	// Variable declaration
	var number int

	// Variable initialization
	fmt.Printf("Enter the number of loop: ")
	fmt.Scan(&number)

	// for loop statement (simple)
	for x := 1; x <= number; x++ {
		fmt.Printf("The number is %v\n", x)
	}

	// for loop statement (custom)
	x := 1
	for x <= number {
		fmt.Printf("The custom looping number is %v\n", x)
		x++
	}

	// for loop statement (with break and continue)
	for x := 0; x <= number; x++ {
		if x == 0 {
			fmt.Printf("The loop index is ignore by continue\n")
			continue
		}
		if x == 1 {
			fmt.Printf("The loop is break\n")
			break
		}
		fmt.Printf("Do somethin..with this index[%v]\n", x)
	}

}
