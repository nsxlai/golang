package main

import (
	"fmt"
)

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	c := Car{
		Name:       "Ferrari",
		Model:      "GTC4",
		Color:      "Red",
		WeightInKg: 1920,
	}

	// Accessing struct fields using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"
	fmt.Println("Car: ", c)

	c2 := new(Car)
	c2.Name = "Tesla"
	c2.Model = "Model 3"
	c2.Color = "Black"
	c2.WeightInKg = 2000

	fmt.Println("Car2:", c2)
}
