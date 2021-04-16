package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int = 8
	var num2 int = 4
	var num3 float32 = 8
	var num4 float32 = 5
	answer1 := num1 + num2
	answer2 := (num1 / num2) + 5
	answer3 := (num3 / num4) + 5 // Add float and int is okay
	// answer4 := num1 / num3  // The types are mismatched and will create error
	answer5 := num1 % num2
	answer6 := math.Pi
	fmt.Printf("answer1 = %d\n", answer1)
	fmt.Printf("answer2 = %d\n", answer2)
	fmt.Printf("answer3 = %g\n", answer3)
	// fmt.Printf("answer4 = %g\n", answer4)
	fmt.Printf("answer5 = %d\n", answer5)
	fmt.Printf("answer6 = %.4g\n", answer6)
}
