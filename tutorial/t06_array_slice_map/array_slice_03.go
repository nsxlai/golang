// source: https://gobyexample.com/arrays
// source: https://gobyexample.com/maps
// source: https://gobyexample.com/range
package main

import "fmt"

func main() {
	// Initialize 5-element array with 0
	var a [5]int
	// a := [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	fmt.Println("b[:3] = ", b[:3])
	fmt.Println("b[3:] = ", b[3:])

	// 2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	fmt.Println()


	//----------------------------------------------
	// Iterating list (index and value)
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num%3 == 0 {
			fmt.Println("index:", i)
		}
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the
	// rune and the second the rune itself.
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Printf("%d %3d %U %q\n", i, c, c, c)
	}
}
