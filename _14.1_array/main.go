package main

import "fmt"

func main() {

	// Arrays declaration and initialization
	var countries = [5]string{
		"Bangladesh", "Pakistan", "Nepal", "Findland", "India",
	}

	// show single output
	var country = countries[0]
	fmt.Println(country)

	// or, show multiple output
	length := len(countries)
	for i := 0; i < length; i++ {
		fmt.Println(countries[i])
	}
}
