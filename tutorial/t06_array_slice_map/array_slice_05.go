package main

import "fmt"

func main() {
	a := []string{"test01", "test02"}
	b := []string{"test03", "test04", "test05"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("a =", a)

	p := []byte{2, 3, 5}
	p = append(p, 7, 11, 13)
	fmt.Println("p =", p)

	d := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	e1 := d[:2]
	e2 := d[2:]
	fmt.Printf("d_q = %q\n", d)
	fmt.Printf("d_t = %t\n", d)
	fmt.Printf("d_T = %T\n", d)
	fmt.Printf("d_s = %s\n", d)
	fmt.Printf("d_s = %b\n", d)
	fmt.Printf("e1 = %q", e1, "e2 = %q", e2)

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}

endofprogram:
	fmt.Println("end of program")
}
