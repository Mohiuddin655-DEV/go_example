package main

import "fmt"

func main() {

	// Arrays declaration
	var countries [2]string

	// Array initialization (single by single)
	fmt.Printf("Enter the country name for index 0 : ")
	fmt.Scan(&countries[0])

	fmt.Printf("Enter the country name for index 1 : ")
	fmt.Scan(&countries[1])

	// show output
	fmt.Println(countries[0])
	fmt.Println(countries[1])
}
