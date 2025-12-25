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

	v, ok := mp["Mia"]
	if ok {
		fmt.Println("Age of Mia is:", v)
	} else {
		fmt.Println("Mia not found")
	}

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

	// Increment
	mx := make(map[string]int)
	mx["counter"] = 0
	fmt.Println("Initial:", mx["counter"])
	mx["counter"]++
	fmt.Println("After Increment:", mx["counter"])
}
