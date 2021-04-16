package main

import "fmt"

func plusTwoNum(a int, b int) int {
	return a + b
}

func plusThreeNum(a, b, c int) int {
	return a + b + c
}

func multipleReturn() (int, int) {
	return 10, 100
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plusTwoNum(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusThreeNum(1, 2, 3) // res is already defined as int, so no := needed here
	fmt.Println("1 + 2 + 3 =", res)

	// Unpacking multiple value to a and b variables.
	a, b := multipleReturn()
	fmt.Println("a =", a, "and b =", b) // auto padding extra " " before and after the string

	c, d := a, b
	fmt.Println("c =", c, "and d =", d)

	c, d = b, a
	fmt.Println("c =", c, "and d =", d)

	// Variadic functions (unspecified nums of function variable input, like *arg in Python)
	sum(1, 2)
	sum(1, 2, 3, 4, 5, 6)

	nums := []int{10, 20, 30, 40, 50}
	sum(nums...)
}
