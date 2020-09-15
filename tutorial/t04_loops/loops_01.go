package main

import "fmt"

func main() {
	// while equivalent loop
	fmt.Println("while equivalent loop")
	x := 3
	for x < 10 {
		fmt.Println(x)
		//x += 1
		x++
	}

	// C style index advancing for loop
	fmt.Println("\nC style index advancing for loop")
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	// break and continue keywords work like Python
	// The follow loop will print numbers NOT divisible by 3, 7, and 9
	fmt.Println("\nfor loop with if-continue condition")
	for x := 0; x <= 50; x++ {
		if x != 0 && x%3 == 0 {
			fmt.Println("Satify condition: ", x)
			continue
		}
		fmt.Printf("%d ", x)
	}
	fmt.Println()

	// while forever loop (while True)
	fmt.Println("\nWhile forever loop")
	for {
		fmt.Println("This loop will run forever if not break, unless use the keyword 'break'")
		break
	}

	// Access array index and value
	fmt.Println("\nEnumerating array index and value in the loop")
	rvar := []string{"This", "is", "a", "test"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	fmt.Println("\nLooping over both index and value")
	for i, j := range rvar {
		fmt.Println(i, j)
	}

	fmt.Println("\nLooping over index only")
	for i := range rvar {
		fmt.Println(i)
	}

	fmt.Println("\nLooping over value only")
	for _, j := range rvar {
		fmt.Println(j)
	}

	// Iterating through a string
	fmt.Println()
	for i, j := range "XabCd" {
		fmt.Printf("index = %d, q = %q, t = %T, b = %b, d = %d\n", i, j, j, j, j)
		fmt.Println(i, j)
	}
}
