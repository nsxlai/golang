// source: https://blog.golang.org/slices-intro
package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10) // Init s with initial length of 5 but capacity of 10
	fmt.Println("Before init: s1 =", s1)

	for i := 1; i < 5; i++ {
		s1[i] = i
	}
	fmt.Println("After init: s1 =", s1)

	// s2 := make([]int{10, 20, 30, 40, 50}, 8, 12)  // Cannot initialize in the make function
	s2 := make([]int, 8, 12)
	s2 = []int{10, 20, 30, 40} // Overwrites the initial make function declaration
	fmt.Println("s2 =", s2, "\nlen(s2) =", len(s2), "cap(s2) =", cap(s2))

	s3 := make([]int, 5, 10)
	s3[0] = 8
	s3[1] = 12
	s3[2] = 16
	fmt.Println("Log1: s3 =", s3)
	fmt.Println("Log1: len(s3) =", len(s3), "cap(s3) =", cap(s3))
	s3[3] = 20
	s3[4] = 24
	// s3[5] = 28  // Set S3[5] (6th element in a length 5 array) will result "panic: runtime error: index out of range"
	fmt.Println("Log2: s3 =", s3)
	fmt.Println("Log2: len(s3) =", len(s3), "cap(s3) =", cap(s3))

	// To extend the array capacity:
	s3 = s3[:6] // Able to set the 6th element
	s3[5] = 28
	fmt.Println("Log3: s3 =", s3)
	fmt.Println("Log3: len(s3) =", len(s3), "cap(s3) =", cap(s3))

	// To extend to array full capacity:
	s3 = s3[:cap(s3)] // Initialize the extra element space with 0
	fmt.Println("Log4: s3 =", s3)
	fmt.Println("Log4: len(s3) =", len(s3), "cap(s3) =", cap(s3))

	// To extend to more than array full capacity:
	s3 = append(s3, 32, 36)
	fmt.Println("Log5: s3 =", s3)
	fmt.Println("Log5: len(s3) =", len(s3), "cap(s3) =", cap(s3))

	s3 = append(s3, 40, 44, 48, 52, 56, 60, 64, 68, 72)           // this will add additional 9 element space to the same array
	fmt.Println("Log6: s3 =", s3)                                 // array capacity doubles every time with "append" function
	fmt.Println("Log6: len(s3) =", len(s3), "cap(s3) =", cap(s3)) // The capacity is now at 40

	s3 = s3[:25] // only reduce the length down to 25, but the capacity is still at 40
	fmt.Println("Log7: s3 =", s3)
	fmt.Println("Log7: len(s3) =", len(s3), "cap(s3) =", cap(s3))

	// Copy s3 to s4; s4 will be a smaller array size than s3 so the copy content will be truncated
	s4 := make([]int, 15, 30)
	copy(s4, s3)
	fmt.Println("Log8: s4 =", s4)
	fmt.Println("Log8: len(s4) =", len(s4), "cap(s4) =", cap(s4))
}
