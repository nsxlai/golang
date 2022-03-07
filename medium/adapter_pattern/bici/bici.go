// source: https://carlos-garcia-13446.medium.com/adapter-pattern-with-golang-2e64598eef48
package bici

import "fmt"

type Bicicleta struct {
	Marca string
	Color string
}

func (b *Bicicleta) Avanzar() {
	fmt.Println("Avanzando!")
}
