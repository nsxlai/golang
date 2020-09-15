package main

import "fmt"

func main() {

	n1, l1 := fullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of chars: %v\n\n", n1, l1)

	n2, l2 := fullNameNakedReturn("Author", "Dent")
	fmt.Printf("Fullname: %v, number of chars: %v\n\n", n2, l2)
}

func fullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func fullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l // change from := to =
	length = len(full) // change from := to =
	return
}
