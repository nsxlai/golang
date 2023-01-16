package main

import (
	"fmt"
	"main/tutorial/t08_multiple_files/model"
)

func main() {
	c := model.Customer{
		Id:   1,
		Name: "TestUser001",
	}
	// c.isEnrolled = true // Error: can not refer to unexported field or method
	// a := model.address{} // Error: can not refer to unexported name

	fmt.Println("TestUser001 = \n", c)

	d := model.Customer{
		Id:   2,
		Name: "TestUser002",
		// Addr: {
		// 	street:  "123 Big Tree Street",
		// 	city:    "Whosville",
		// 	state:   "CA",
		// 	country: "USA",
		// },
		// isEnroll: true,
	}

	fmt.Println("TestUser002 = \n", d)
}
