package main

import "fmt"

func main() {

	// arrays declaration
	var countries []string

	// array size declaration
	var size int

	// array size initialization
	fmt.Printf("How many countries : ")
	fmt.Scan(&size)

	// Array initialization

	for i := 0; i < size; i++ {

		// array temporary element
		var country string

		// user input
		fmt.Printf("Enter the country  : ")
		fmt.Scan(&country)

		// array initialize
		countries = append(countries, country)
	}

	// show output
	fmt.Println(countries)
}
