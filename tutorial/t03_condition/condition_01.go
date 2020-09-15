package main

import "fmt"

func main() {
	x := 5
	y := 6.6
	val1 := x <= 5
	fmt.Printf("val1 = %t\n", val1)
	fmt.Printf("val2 = %t\n", float64(x)+1.5 == y)

	val3 := 5 < 7 || 7 > 9          // || = or
	fmt.Printf("val3 = %t\n", val3) //true

	val4 := !(5 < 7) // false
	fmt.Printf("val4 = %t\n", val4)
	fmt.Printf("expression1: %t\n", true || false && true)
	fmt.Printf("expression2: %t\n", 5 < 10 || 50 < 20)

	fmt.Println("Before if")
	if !val4 {
		fmt.Println("inside if")
	} else if val3 {
		fmt.Println("inside else if")
	} else {
		fmt.Println("inside else")
	}
}
