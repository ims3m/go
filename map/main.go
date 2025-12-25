package main

import "fmt"

func main() {
	var mp = map[string]int{
		"Sam":   25,
		"Steve": 30,
		"Lucy":  22,
		"Mia":   28,
	}

	fmt.Println(mp)
	fmt.Println("Age of Steve:", mp["Steve"])
	fmt.Println("Length of Map is:", len(mp))

	var mt map[string]string

	mt = make(map[string]string)

	mt["USA"] = "Washington D.C."
	mt["France"] = "Paris"
	mt["Japan"] = "Tokyo"

	fmt.Println(mt)

	delete(mt, "France")
	fmt.Println("Map After Deletion: ", mt)

	// Iterating over Map
	for k, v := range mp {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
}
