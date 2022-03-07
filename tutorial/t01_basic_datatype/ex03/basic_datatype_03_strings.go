package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = "test"
	var s2 string = "a"
	s3 := "The weather is beautiful"
	s4 := string("poster")

	// array
	var s5 [6]string
	s5[0] = "t"
	s5[1] = "e"
	s5[2] = "s"
	s5[3] = "t"
	s5[4] = "0"
	s5[5] = "1"

	// same as:
	s51 := [6]string{"t", "e", "s", "t", "0", "2"}
	s52 := [...]string{"t", "e", "s", "t", "0", "3"}

	// slice
	s6 := make([]string, 10)
	s6[0] = "t"
	s6[1] = "e"
	s6[2] = "s"
	s6[3] = "t"
	s6[4] = "0"
	s6[5] = "4"

	// Three way of declaring string slices
	var s7 = []string{"test01", "test02"}
	s8 := []string{"test03", "test04", "test05"}
	s9 := make([]string, 10)
	// s9 = "abc"  wouldn't work as declared string array (but will work as regular string)
	s9[0] = "a"
	s9[1] = "b"
	s9[2] = "c"

	fmt.Printf("s1 = %v\n", s1)
	fmt.Println("Upper case s1: ", strings.ToUpper(s1))
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("s3 = %q\n", s3)
	fmt.Printf("s3 = %s\n", s3)
	fmt.Printf("s3 = %v\n", s3)
	fmt.Println("Title s3: ", strings.Title(s3)) // capitalize all the word in the string
	fmt.Printf("len(s3) = %d\n", len(s3))        // array has no cap function supported (only slice has cap function)
	fmt.Printf("s3[0] = %v\n", s3[0])
	fmt.Printf("s3[0] = %s\n", s3[0])
	fmt.Printf("s3[0] = %q\n", s3[0])
	fmt.Printf("s4 = %v\n", s4)
	fmt.Printf("s5 = %v\n", s5)
	fmt.Printf("s5 = %s\n", s5)
	fmt.Printf("s5 = %q\n", s5)
	fmt.Printf("s51 = %q\n", s51)
	fmt.Printf("s52 = %v\n", s52)
	// fmt.Println("Joining s5 and s4: ", append(s5, s4...))  // append does not work on array
	fmt.Printf("cap(s6) = %d; len(s6) = %d\n", cap(s6), len(s6))
	s6 = append(s6, s6...)
	fmt.Println("Joining s6 and s6: ", s6) // both appending objects have to be slice

	fmt.Println("(before append) s7 =", s7)
	s7 = append(s7, s8...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("(after append) s7 =", s7)

	fmt.Println("s9 =", s9)

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("Equal?", lValue == uValue)
	fmt.Println("Equal Non-Case Sensitive?",
		strings.EqualFold(lValue, uValue)) // case folding

	str1 := "An implicitly typed string"
	str2 := "An explicitly typed string"
	fmt.Println("str1 Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("str2 Contains exp?", strings.Contains(str2, "exp"))

	// zero argument
	fmt.Println("empty joinstr: ", joinstr())

	// multiple arguments
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println(joinstr("G", "E", "E", "k", "S"))
}

// Variadic function to join strings
func joinstr(element ...string) string {
	return strings.Join(element, "-")
}
