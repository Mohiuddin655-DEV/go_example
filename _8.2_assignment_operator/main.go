package main

import "fmt"

func main() {

	// Types of  assignment operators => =, +=, -+, *=, /=, %=

	value := 0

	value = 5
	fmt.Println("Assignment operator (=)  : ", value)

	value += 5
	fmt.Println("Assignment operator (+=) : ", value)

	value -= 5
	fmt.Println("Assignment operator (-=) : ", value)

	value *= 5
	fmt.Println("Assignment operator (*=) : ", value)

	value /= 5
	fmt.Println("Assignment operator (/=) : ", value)

	value %= 5
	fmt.Println("Assignment operator (%=) : ", value)
}
