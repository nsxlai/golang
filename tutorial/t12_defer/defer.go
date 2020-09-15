package main

import "fmt"

func main() {
	// will execute the defer function after everything else has executed in the function
	// each defer is put in a FILO stack
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")
	fmt.Println("Do something with the file!")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undeferred Statement")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Value of x after incrementing:", x)
}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
