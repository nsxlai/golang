package model

// Unexported struct (only accessible inside package `model`)
type address struct {
	street, city, state, country string
	zipCode                      int
}
