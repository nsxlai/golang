package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)  // will parse the user input and find the first space

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // TrimSpace remove the \n at the end (remove proceeding white space char as well)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", f)
	}
}
