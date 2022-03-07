// source: https://blog.golang.org/slices-intro#:~:text=Unlike%20an%20array%20type%2C%20a%20slice%20type%20has%20no%20specified%20length.&text=where%20T%20stands%20for%20the,that%20refers%20to%20that%20array.
package main

import (
	"fmt"
)

func main() {
	// Array declaration 1
	var a [4]int // declare variable a as an array of 4 integers
	a[0] = 1
	i := a[0]
	fmt.Println("a = ", a) // When initialized, will fill all the array element with 0
	fmt.Println("i = ", i)

	// Array declaration 2
	b := [2]string{"alpha", "centuri"}
	fmt.Println("b = ", b)

	// If the elements are too many, ... can summarize the count
	c := [...]string{"aaa", "aab", "aac", "aad", "aba", "abb", "abc", "abd"}
	fmt.Println("c = ", c)
	fmt.Println("len(c) = ", len(c))

	// Unlike an array type, a slice type has no specified length
	// Slice declaration 1
	slice01 := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("slice01 = ", slice01)

	// Slice delcaraion 2
	var slice02 []byte
	slice02 = make([]byte, 5) // Len(slice02) == 5
	fmt.Println("slice02 = ", slice02)
	slice02[0] = 'h'
	slice02[1] = 'e'
	slice02[2] = 'l'
	slice02[3] = 'l'
	slice02[4] = 'o'
	fmt.Println("slice02 = ", slice02)

	// Slice declaration 3
	slice03 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println("slice03 = ", slice03)

	// slicing
	fmt.Println("slice03[1:4] = ", slice03[1:4])
	temp := []byte{'o', 'l', 'a'}
	fmt.Printf("Slice Test = %t\n", temp)
	fmt.Println("Slice Test = ", temp)
}
