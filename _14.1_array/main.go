package main

import "fmt"

func main() {

	// Arrays declaration and initialization
	var countries = [5]string{
		"Bangladesh", "Pakistan", "Nepal", "Findland", "India",
	}

	// show output
	fmt.Println(countries[0])
	fmt.Println(countries[3])

	// or, show output
	length := len(countries)
	for i := 0; i < length; i++ {
		fmt.Println(countries[i])
	}
}
