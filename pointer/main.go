package main

import "fmt"

func chDelta(v *int) {
	*v = 10
}

type Person struct {
	name   string
	health int
}

func playerTakesDamage(p *Person) {
	p.health -= 10
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

	mx := Person{
		name:   "Hero",
		health: 100,
	}

	// Pointer is a 8 byte long integer
	fmt.Printf("Before taking damage: %d\n", mx.health)
	playerTakesDamage(&mx)
	fmt.Printf("After taking damage : %d\n", mx.health)

}
