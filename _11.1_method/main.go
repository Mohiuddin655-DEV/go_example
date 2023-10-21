package main

import (
	"fmt"
)

func main() {

	// call simple method
	sayHelloWorld()

	// call single parameterized method
	printMessage("Hey, I'm here...!")
	printMessage("I'm from Bangladesh.")

	// call multiple parameterized method
	sayIntroduction("Md. Mohi-Uddin", 23, "Bangladesh")
	sayIntroduction("Mohammad", 26, "Soudi Arabia")

	// call return type method
	var intro = getMyIntroduction()

	// print the introduction
	printMessage(intro)
}

// void method
func sayHelloWorld() {
	fmt.Printf("Hello World!\n")
}

// single parameterized method
func printMessage(message string) {
	fmt.Printf("%v\n", message)
}

// multiple parameterized method
func sayIntroduction(name string, age int, country string) {
	fmt.Printf("Hey, I'm %v. I'm from %v. I'm currently %v.\n", name, country, age)
}

// return type method
func getMyIntroduction() string {
	var intro = "Hey, I'm Md. Mohi-Uddin. I'm from Bangladesh. I'm currently 23."
	return intro
}
