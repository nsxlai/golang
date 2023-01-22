// source: https://medium.com/towardsdev/golang-solid-principles-e7641dee90b0
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	// shape
	volume() float64
}

type square struct {
	sideLen float64
}

func (s square) area() float64 {
	return math.Pow(s.sideLen, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.sideLen, 3)
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func volumeSum(objects ...object) float64 {
	var sum float64
	for _, s := range objects {
		sum += s.volume()
	}
	return sum
}

func main() {
	s1 := square{sideLen: 2}
	s2 := square{sideLen: 5}
	c1 := cube{square{sideLen: 3}}
	c2 := cube{square{sideLen: 4}}

	sumArea := areaSum(s1, s2)
	sumVolume := volumeSum(c1, c2)

	fmt.Println("Sum of Area: ", sumArea)
	fmt.Println("Sum of Volume: ", sumVolume)
}
