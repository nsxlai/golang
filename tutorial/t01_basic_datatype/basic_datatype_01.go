package main

import "fmt"

func main() {
	// var type:
	// int8, int16, int32, int64,
	// uint8, uint16, uint32, uint64,
	// float32, float64,
	// complex64, complex128,
	// byte (=uint8), rune (=int32), uint (32 or 64 bits), int (either 32 or 64 bits, same as uint)
	// Once type is declared, cannot change at runtime

	// explicit type declaration
	var number uint16 = 260
	var number1 = 2
	var number2 = 3.5
	var number2x float32 = 2.5

	// implicit type declaration
	number3 := 6
	b1 := true

	// explicit constant typing declaration
	const sInt int = 500

	// implicit constant typing declaration (no := required)
	const s = 1000

	var b2 bool
	var number4 uint64
	var testName string = "Hello World!"
	test1 := "A"
	test2 := 'A'

	number += 100
	fmt.Println(testName, number)
	fmt.Printf("number 1: %T\n", number1)
	fmt.Printf("number 2: %T\n", number2)   // Default float is float64
	fmt.Printf("number 2x: %T\n", number2x) // unless specified to float32
	fmt.Printf("number 3: %T\n", number3)
	fmt.Printf("convert to different data type: %T\n", int16(number2x))
	fmt.Printf("Constant type: %T\n", s)
	fmt.Printf("b1: %T\n", b1) // %T will show the data type
	fmt.Printf("b2: %T\n", b2)
	fmt.Printf("b1: %t\n", b1) // %t will represent the boolean value
	fmt.Println("b2: ", b2)
	fmt.Println("!true = ", !true)
	fmt.Printf("number4: %T\n", number4)
	fmt.Printf("number4: %v\n", number4)
	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T",
		number1, number2, number2x, b1, testName)
	fmt.Println("mystring: ", myString)

	fmt.Printf("3025 octal: %o\n", 3025)
	fmt.Printf("3025 hexdecimal (non-cap): %x\n", 3025)
	fmt.Printf("3025 hexdecimal (cap): %X\n", 3025)
	fmt.Printf("3025 binary: %b\n", 3025)

	fmt.Printf("String example: %s\n", "test")
	fmt.Printf("String example: %q\n", "test")
	fmt.Println()
	fmt.Println("Example" + " combining" + " strings\n")

	fmt.Printf("[test1] 1 char in string: %T\n", test1)
	fmt.Printf("[test1] 1 char in string: %b\n", test1)
	fmt.Printf("[test1] 1 char in string: %s\n", test1)
	fmt.Printf("[test1] 1 char in string: %q\n", test1)
	fmt.Printf("[test1] 1 char in string: %v\n", test1)

	fmt.Printf("[test2] byte: %T\n", test2)
	fmt.Printf("[test2] byte: %b\n", test2)
	fmt.Printf("[test2] byte: %s\n", test2)
	fmt.Printf("[test2] byte: %q\n", test2)
	fmt.Printf("[test2] byte: %v\n", test2)

	fmt.Printf("Float number: %f\n", 3.1415927)
	fmt.Printf("Float number: %010f\n", 3.1415927)
	fmt.Printf("Float number: %9.5f\n", 3.1415927)
	fmt.Printf("Float number: %.4f\n", 3.1415927)

	fmt.Printf("INT padding: %09d\n", number)
	fmt.Printf("INT padding: %-4d\n", number) // left justify

	fmt.Printf("test1 type: %T\n", test1)
	fmt.Printf("Unicode for %s: %T\n", test1, test1)

	var out string = fmt.Sprintf("Number: %07d", 45)
	fmt.Println(out)
}
