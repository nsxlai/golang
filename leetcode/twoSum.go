package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 17

	i1, i2 := twoSum(nums, target)
	fmt.Printf("twoSum result: %v, %v", i1, i2)
}

func twoSum(nums []int, target int) (int, int) {
	var complement int
	for i := range nums {
		complement = target - nums[i]
		index, ok := contains(nums[i:], complement)
		if ok {
			return i, index
		}
	}
	return -1, -1
}

func contains(s []int, e int) (int, bool) {
	for index, a := range s {
		if a == e {
			return index, true
		}
	}
	return -1, false
}
