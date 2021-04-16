package main

import (
	"fmt"
)

func main() {

	name := "test01"
	show := func(name string) string {
		out := fmt.Sprintf("This is %s", name)
		return out
	}

	pshow := func(name string) {
		fmt.Println("This is also", name)
	}

	
	fmt.Println(show(name))
	pshow(name)
}
