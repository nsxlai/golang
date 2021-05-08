package main

import "fmt"

func main() {
	// Example is sourced from
	// https://betterprogramming.pub/go-4f365468dbd5
	fmt.Print("Enter something: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Input  " + first + "\n")
}
