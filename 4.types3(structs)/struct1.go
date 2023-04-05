//* Crear un struct de tipo persona, debe almacenar:
//* nombre
//* apellido
//* comidas favoritas
//* Crear dos personas, imprimir el nombre y las comidas.

package main

import (
    "fmt"
)

//creo struct de tipo Persona 
type Persona struct {
	nombre string
	apellido string
	comidasFavoritas []string
}


func main() {

	//creo las personas
	p1 := Persona{
		nombre: "georgiana",
        apellido: "garino",
        comidasFavoritas: []string {"milanesas","pizzas","chinchulin"},
	}
	p2 := Persona{
		nombre: "damian",
        apellido: "marquez",
        comidasFavoritas:  []string {"empanadas","hamburguesa","papas"},
	}

	//imprimir el nombre y las comudas

	fmt.Println("el nombre de la primer persona es: ", p1.nombre, "\n las comidas favoritas de la primer persona es: ", p1.comidasFavoritas)
	fmt.Println("el nombre de la segunda persona es: ", p2.nombre, "\n las comidas favoritas de la segunda persona es: ", p2.comidasFavoritas)

}

