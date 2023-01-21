// source: https://waclawthedev.medium.com/loop-pointer-in-go-golang-small-trap-with-big-consequences-220e4fab15e0
package main

import "fmt"

func main() {

	sourceArray := []string{"first", "second", "third"}
	var destArray1 []*string

	// The wrong way
	for _, element := range sourceArray {
		destArray1 = append(destArray1, &element)
	}

	// Trying to print elements of destArray. The result is not as intended
	for _, element := range destArray1 {
		println(*element)
	}
	fmt.Println()

	//--------------------------------------------------------------
	// To correct the bug:
	// First print the sourceArray content
	var destArray2 []*string

	for i, elem := range sourceArray {
		fmt.Printf("Index: %d, element: %s, pointer: %p \n", i, elem, &elem)
	}

	// Now we are using index and access to element by index
	for i, _ := range sourceArray {
		destArray2 = append(destArray2, &sourceArray[i])
	}
	fmt.Println()

	// Trying to print elements of destArray
	for _, element := range destArray2 {
		println(*element)
	}
}
