// source: https://gobyexample.com/arrays
// source: https://gobyexample.com/maps
// source: https://gobyexample.com/range
package main

import "fmt"

func main() {
	//----------------------------------------------
	// Maps (Python dictionary)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	// fmt.Println("_:", _)   // Unable to print _ value, unlike Python
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	fmt.Println()

	//----------------------------------------------
	// Iterating a map (index and value)
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "carrot"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
}
