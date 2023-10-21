package main

import "fmt"

type Person struct {
	name    string
	age     int
	country string
}

// struct wise method
func (person *Person) updatePersonName(name string) {
	person.name = name
}

// struct wise method
func (person *Person) updatePersonAge(age int) {
	person.age = age
}

// struct wise method
func (person *Person) updatePersonCountry(country string) {
	person.country = country
}

// struct wise method
func (person *Person) showPersonInfo() {
	fmt.Printf("Name : %v, Age : %v, Country : %v\n", person.name, person.age, person.country)
}

func main() {

	person1 := Person{"Md. Mohi-Uddin", 23, "Bangladesh"}

	person2 := Person{
		"Omie Talukdar",
		24,
		"Bangladesh",
	}

	person3 := Person{
		"Mohammad",
		26,
		"Soudi Arabia",
	}

	// call to create biograpy for both person data
	createBio(person1)
	createBio(person2)
	createBio(person3)

	// person 1 data update operation
	person1.updatePersonName("Mr. Developer")
	person1.updatePersonAge(24)
	person1.updatePersonCountry("Palestine")

	// show person1 updated data
	person1.showPersonInfo()
}

func createBio(person Person) {

	var name = person.name
	var age = person.age
	var country = person.country

	fmt.Printf("My name is %v. I'm from %v. I'm currently %v\n", name, country, age)
}
