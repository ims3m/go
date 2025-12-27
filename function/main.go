package main

import "fmt"

// func (r reciever) identifier (p parameter(s)) (return(s)) { ... }
func main() {

	foo()

	bar("Gopher")

	v := test("Test Value")
	fmt.Println(v)

	// Defer
	defer printAtLast()

	n, a := f("Sam", 24)
	fmt.Println(n, a)

	v1, v2, v3 := ax(3, 4)
	fmt.Println(v1, v2, v3)

	sl := []int{1, 4, 5, 6}
	op := unFunc(sl...)

	fmt.Println(op)

}

// No params, No return
func foo() {
	fmt.Println("Func Foo")
}

func bar(name string) {
	fmt.Println("Hello", name)
}

func test(s string) string {
	return fmt.Sprint("Value is: ", s)
}

func f(name string, age int) (string, int) {
	return fmt.Sprint("My name is ", name), age
}

func ax(a int, b int) (int, int, int) {
	var c int
	c = a * b
	return a, b, c
}

// Variadic Parameters
func variadicFunc(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	fmt.Println("Sum: ", sum)
	return sum
}

// Variadic Function (unfurling)
func unFunc(num ...int) int {
	val := 0

	for _, v := range num {
		val += v
	}
	return val
}

func printAtLast() string {
	fmt.Printf("This will be printed at last\n")
	return "End of Function"
}

type Cat struct {
	Name string
}

func (c Cat) String() string {
	fmt.Println("Cat Name: ", c.Name)
	return c.Name
}
