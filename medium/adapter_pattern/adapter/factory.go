package adapter

import (
	// "github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/02-adapter/auto"
	// "github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/02-adapter/bici"
	auto "../auto"
	bici "../bici"
)

func Factory(s string) Adapter {
	switch s {
	case "automovil":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bicicleta":
		d := bici.Bicicleta{}
		return &BicicletaAdapter{&d}
	}

	return nil
}
