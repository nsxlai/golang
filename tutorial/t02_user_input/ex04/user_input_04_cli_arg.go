package main

import (
	"fmt"
	"os"
)

func main() {
	// Example is sourced from
	// https://betterprogramming.pub/go-4f365468dbd5
	// At CLI:
	// go run user_input_04_cli_arg.go test_001
	arg := os.Args[1]
	fmt.Print("Argument: " + arg)
}
