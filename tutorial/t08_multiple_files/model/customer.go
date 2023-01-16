package model

type Customer struct { // exported struct type
	Id       int     // exported field
	Name     string  // exported field
	addr     address // unexported field (only accessible inside package `model`)
	isEnroll bool    // unexported field (only accessible inside package `model`)
}
