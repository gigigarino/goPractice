package main
//* Crear struct tipo vehículo, con los campos:
//* puertas
//* color
//* Crear 2 nuevos tipos:
//* sedan
//* camion
//* Embeber vehículo en ambos.
//* Para camión agregar el campo "cuatro Ruedas" bool.
////* Para sedan agregar el campo "deLujo" bool.
//* Con un composite literal, crear el valor de tipo camión y asignar los valores a los campos.
//* Con un composite literal, crear el valor de tipo sedan y asignar los valores a los campos.

import (
    "fmt"
)

//crear el struct vehiculo
//crear 2 tipos de vehiculos
//* sedan
//* camion
type Vehiculo struct {
	puertas int
	color string
}


type Sedan struct {
	Vehiculo
	deLujo bool
}

type Camion struct {
	Vehiculo
	cuatroRuedas bool
}

//* Embeber vehículo en ambos.
func main() {
	c := Camion{
		Vehiculo: Vehiculo{
			puertas: 4,
            color: "rojo",
		},
		cuatroRuedas: true,
	}

	s := Sedan{
		Vehiculo: Vehiculo{
			puertas: 5,
			color: "azul",
		},
		deLujo: true,
	}

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(c.puertas)
	fmt.Println(s.puertas)
}