package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors [3]string // cannot mix different types in the same array
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))

	// can change array contents:
	colors[0] = "Yellow"
	fmt.Println(colors)

	// However array is unable to add or remove elements at runtime
	// for those features, should use slices instead of array
	var col = []string{"Red", "Green", "Blue"}
	fmt.Println("slice: ", col)

	// append function is built to work with slice
	col = append(col, "Purple")
	fmt.Println("slice after append: ", col)

	// remove the first element
	col = append(col[1:len(col)])
	fmt.Println("remove the first element: ", col)

	col = append(col[:len(col)-1])
	fmt.Println("remove the last element: ", col)

	num := make([]int, 5, 5)
	num[0] = 1
	num[1] = 7
	num[2] = 32
	num[3] = 12
	num[4] = 156
	fmt.Println(num)

	num = append(num, 235)
	fmt.Println("Capacity after adding one more element: ", cap(num))

	// able to sort slice as well
	sort.Ints(num)
	fmt.Println("Sorted numbers:", num)
}
