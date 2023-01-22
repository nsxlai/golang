// source: https://medium.com/towardsdev/golang-solid-principles-e7641dee90b0
package main

import (
	"math"
)

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type Square struct {
	sideLen float32
}

func (s Square) area() float32 {
	return s.sideLen * s.sideLen
}
