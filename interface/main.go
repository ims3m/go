package main

import "fmt"

type Person struct {
	name string
}

type Employee struct {
	Person
	id int
}

func main() {
	fmt.Println("This is the main package.")
}
