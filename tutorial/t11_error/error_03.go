package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("filename.txt")
	fmt.Printf("err type: %T\nerr value: %v\n", err, err)
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("My error string") // Create custom error function
	fmt.Println(myError)

	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true}
	attended, ok := attendance["Mike"]
	fmt.Printf("ok: %t\n", ok)
	if ok {
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info for Mike")
	}

	attended, ok = attendance["Lisa"]
	fmt.Printf("ok: %t\n", ok)
	if ok {
		fmt.Println("Lisa attended?", attended)
	} else {
		fmt.Println("No info for Lisa")
	}
}
