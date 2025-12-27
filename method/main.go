package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Speak() {
	fmt.Println("Hello! my name is", p.name)
}

func main() {

	p1 := Person{name: "Sam"}

	p2 := Person{name: "Alex"}

	p2.Speak()
	p1.Speak()

}
