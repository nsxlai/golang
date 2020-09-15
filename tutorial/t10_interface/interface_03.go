// https://www.callicoder.com/golang-interfaces/
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct type `Rectangle` - implements the `Shape` interface by implementing all its methods.
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Struct type `Circle` - implements the `Shape` interface by implementing all its methods.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// Interface types can also be used as fields
type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (drawing MyDrawing) Area() float64 {
	totalArea := 0.0
	for _, s := range drawing.shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {
	drawing := MyDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{3, 5},
			Rectangle{4, 7},
		},
		bgColor: "red",
		fgColor: "white",
	}

	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())
}
