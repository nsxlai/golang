// source: https://medium.com/towardsdev/golang-solid-principles-e7641dee90b0
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

type triangle struct {
	height, base float32
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}

type shape interface {
	area() float32
}

type calculator struct {
	total float32
}

func (c *calculator) sumAreas(shapes ...shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{sideLen: 2}
	t := triangle{height: 10, base: 5}

	calculator := calculator{total: 0}

	fmt.Println("The total area is: ", calculator.sumAreas(c, s, t))
}
