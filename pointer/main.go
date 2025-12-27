package main

import "fmt"

func chDelta(v *int) {
	*v = 10
}

func main() {

	x := 4
	// Value of x
	fmt.Println(x)
	// Memory address of x
	fmt.Println(&x)

	/*
		(*int) is a pointer to an int
	*/

	y := &x
	fmt.Println(y)
	// Dereferencing the pointer to get the value
	fmt.Println(*y)

	/*
		fmt.Println(*&x)
	*/

	*y = 10
	fmt.Println(x)

	b := 42

	fmt.Println(b)
	chDelta(&b)
	fmt.Println(b)

}
