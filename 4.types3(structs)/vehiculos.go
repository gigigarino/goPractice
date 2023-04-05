package main

import (
	"fmt"
)

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	deLujo bool
}

func main() {

	c := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "rojo",
		},
		cuatroRuedas: true,
	}

	s := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "verde",
		},
		deLujo: true,
	}

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(c.puertas)
	fmt.Println(s.puertas)
}