package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	PersonInfo Person
	JobTile    string
}

func main() {

	s1 := Person{
		Name: "Sam",
		Age:  25,
	}

	s2 := Employee{
		PersonInfo: Person{
			Name: "Steve",
			Age:  30,
		},
		JobTile: "Developer",
	}

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%s %s %s\n", s1.Name, s2.PersonInfo.Name, s2.JobTile)
}
