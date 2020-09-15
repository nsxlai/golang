package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog"
	aNumber := 42
	isTrue := true

	stringLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String length:", stringLength)
	}

	stringLength2, _ := fmt.Println(str1)
	fmt.Println("string length:", stringLength2)

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of aNumber (float) = %.2f\n", float64(aNumber))
	fmt.Printf("Value of isTrue = %v\n", isTrue)
	fmt.Printf("Value of str1 = %v\n", str1)

	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
