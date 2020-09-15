// https://gobyexample.com/pointers
package main

import "fmt"

func zeroval(ival int) {
	ival = 10
	fmt.Println("ival:", ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
	fmt.Println("iptr:", iptr)
}

func main() {
	i := 1
	fmt.Println("initial value:", i)
	fmt.Println("initial pinter:", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer i:", &i)
	fmt.Println("final value i:", i)

	j := i
	fmt.Println("pointer j:", &j)
	fmt.Println("final value j:", j)

	fmt.Println("Change value of j")
	j = 20
	fmt.Println("pointer i:", &i)
	fmt.Println("final value i:", i)
	fmt.Println("pointer j:", &j)
	fmt.Println("final value j:", j)

}
