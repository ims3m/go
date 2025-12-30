package main

import "fmt"

type Num interface {
	int | float64
}

/*
func getSum[T int | float64](a T, b T) T {
	return a + b
}
*/

func getSum[T int | float64](a T, b T) T {
	return a + b
}

func main() {

	a := 10
	b := 20

	v := getSum(a, b)
	fmt.Println(v)

	x := 10.5
	y := 20.3
	z := getSum(x, y)
	fmt.Println(z)
}
