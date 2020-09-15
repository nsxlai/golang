package main

import (
	"fmt"
)

type Dog1 struct {
	Breed  string
	Weight int
}

type Dog2 struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog2) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog2) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}

func (d *Dog2) SpeakThreeTimesPt() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {

	poodle1 := Dog1{"Poodle1", 34}
	fmt.Println(poodle1)
	fmt.Printf("%+v\n", poodle1) // complete type
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle1.Breed, poodle1.Weight)

	poodle2 := Dog2{"Poodle", 37, "Woof"}
	fmt.Println(poodle2)
	poodle2.Speak()

	poodle2.Sound = "Arf"
	poodle2.Speak()

	poodle2.SpeakThreeTimes()
	poodle2.SpeakThreeTimes()

	fmt.Println()
	poodle2.SpeakThreeTimesPt()
	poodle2.SpeakThreeTimesPt()
}
