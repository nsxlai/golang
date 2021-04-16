// https://medium.com/better-programming/easy-guide-to-unit-testing-in-golang-4fc1e9d96679
package main

import (
	"fmt"
)

func main() {

	op := ops.GetKeyOperator()

	key := op.Generate(2, 3)

	a, b, _ := op.Degenerate(key)

	fmt.Printf("key=%v, a=%v, b=%v", key, a, b)

}
