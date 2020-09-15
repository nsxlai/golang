package main

import (
	"fmt"
	"sort"
)

func main() {
	// Must initialize complex objects before adding values
	// Declarations without make() can cause a panic
	//     var m map[string]int
	//     m["key"] = 42
	//     fmt.Println(m)
	//
	//     // panic: assignment to entry in nil map
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)

	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println("CA = ", california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted keys", keys)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
