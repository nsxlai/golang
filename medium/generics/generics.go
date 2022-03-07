// source: https://levelup.gitconnected.com/generics-in-go-viva-la-revolution-e27898bf5495
//
// To install constraints import (not included originally):
// PS C:\Users\nsxla\PycharmProjects\golang> go get github.com/golang/go/tree/master/src/constraints
// go: downloading github.com/golang/go v0.0.0-20211205125044-ecf6b52b7f4b
// go get: module github.com/golang/go@upgrade found (v0.0.0-20211205125044-ecf6b52b7f4b), but does not contain package github.com/golang/go/tree/master/src/constraints
// 
// go github repo location: https://github.com/golang/go
//
package main

import (
	"constraints"
	"fmt"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Transform[S Number, T Number](input []S) []T {
	output := make([]T, 0, len(input))
	for _, v := range input {
		output = append(output, T(v))
	}
	return output
}

func main() {
	fmt.Printf("%#v", Transform[int, float64]([]int{1, 2, 3, 6}))
}