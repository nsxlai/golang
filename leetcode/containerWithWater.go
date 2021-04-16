package main

import "fmt"

func main() {

	testList := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Max Area =", maxArea(testList))
}

func maxArea(height []int) int {
	maxarea, s := 0, 0
	e := len(height) - 1

	for s < e {
		area := min(height[s], height[e]) * (e - s)
		maxarea = max(area, maxarea)

		if height[s] < height[e] {
			s++
		} else {
			e--
		}
	}
	return maxarea
}

func min(val1, val2 int) int {
	if val1 >= val2 {
		return val2
	} else {
		return val1
	}
}

func max(val1, val2 int) int {
	if val1 <= val2 {
		return val2
	} else {
		return val1
	}
}
