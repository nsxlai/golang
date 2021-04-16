package main

import "fmt"

func main() {
	var arr [50]int
	slc := make([]int, 50)

	for i := 0; i < 50; i++ {
		arr[i] = i
	}
	fmt.Println("arr = ", arr)

	for i := 0; i < 50; i++ {
		slc[i] = i
	}
	fmt.Println("slc = ", slc)
}
