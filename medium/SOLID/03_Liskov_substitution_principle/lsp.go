// source: https://medium.com/towardsdev/golang-solid-principles-e7641dee90b0
package main

import "fmt"

type transport interface {
	getName() string
}

// base entity type
type vehicle struct {
	name string
}

func (c vehicle) getName() string {
	return c.name
}

// subtype car that is composed with vehicle
type car struct {
	vehicle
	wheel int
	gates int
}

type motocycle struct {
	vehicle
	wheel int
}

type printer struct{}

func (printer) printTransportName(p transport) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	v := vehicle{name: "Ford"}

	c := car{
		vehicle: vehicle{name: "F-150 Lightning"},
		wheel:   4,
		gates:   2,
	}

	m := motocycle{
		vehicle: vehicle{name: "Livewire"},
		wheel:   2,
	}

	p := printer{}
	p.printTransportName(v)
	p.printTransportName(c)
	p.printTransportName(m)
}
