//Con el código del ejercicio anterior:
//* Almacenar los valores de tipo persona en un map,
//* con el key del apellido.
//* Acceder a cada valor en el map.
//* Imprimir los valores del slice.

package main

import (
	"fmt"
)

type Persona struct {
	nombre string
	apellido string
	comidasFavoritas []string
}

func main() {
	p1 := Persona{
		nombre: "georgiana",
        apellido: "garino",
        comidasFavoritas: []string{"pizza", "empanadas", "mozzarella", "jamon"},
	}

	p2 := Persona{
		nombre: "matias",
        apellido: "garino",
        comidasFavoritas: []string{"hamburguesas", "carne", "mozzarella", "jamon", "veggies"},
	}

	//hacer el map
	m := map [string] Persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	k := 0
	fmt.Println("--------------------------------")
	for i, v := range m {
        fmt.Println("A", v.nombre, i, "sus comidas favoritas son:")
		for j, comida := range v.comidasFavoritas {
			k = j + 1
			fmt.Println(k, comida)
		}
		fmt.Println("------")
    }
}