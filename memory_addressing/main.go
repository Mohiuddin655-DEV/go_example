package main

import "fmt"

func main() {

	// pointer value roles => &, *
	// 1. & indicate the pointer value
	// 2. * indicate the non-pointer value

	// variable declaration and initialization
	x := 10

	// pointer value (&)
	var p *int
	p = &x

	// non-pointer value (*)
	var q = *p

	// customize memory addressing value
	*p = *p - 1

	fmt.Printf("Simple value is           : %v\n", x)
	fmt.Printf("Pointer value is          : %v\n", p)
	fmt.Printf("Non-Pointer value is      : %v\n", q)
	fmt.Printf("Customized value is       : %v\n", x)

	// change value by reference
	changeValueByReference(&x)
	fmt.Printf("Change value by reference : %v\n", x)

}

func changeValueByReference(reference *int) {
	*reference = 80
}
