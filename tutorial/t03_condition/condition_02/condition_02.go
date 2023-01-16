package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var result string
	var maxInt int = 10

	var y int
	rand.Seed(time.Now().Unix()) // the random package in Go is not as robust as Python. It will create the same random sequence every time unless Seed is changed

	for i := 0; i < maxInt; i++ {
		y = rand.Intn(maxInt)
		fmt.Println("y: ", y)
	}

	// x is only local to the condition statement and will
	// be garbage-collected after the condition is completed
	if x := rand.Intn(maxInt); 5 < x && x <= maxInt {
		result = fmt.Sprintf("Between 5 and %v", maxInt)
	} else if x <= 5 {
		result = "Between 0 and 5"
	} else {
		result = fmt.Sprintf("Outside the range of 0 to %v", maxInt)
	}
	// fmt.Println("x:", x)   // x is undefined at this point
	fmt.Println("Result:", result)

}
