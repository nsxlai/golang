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

type shape interface {
	area() float32
}

type outPrinter struct{}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("The area is: %f", s.area())
}

func main() {
	c := circle{radius: 5}
	c.area()

	s := square{sideLen: 2}
	s.area()

	out := outPrinter{}
	fmt.Println("Circle: ", out.toText(c))
	fmt.Println("Square: ", out.toText(s))
}
