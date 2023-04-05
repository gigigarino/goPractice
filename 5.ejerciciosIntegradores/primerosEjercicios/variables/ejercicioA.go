//a) Crea una variable "nombre" de tipo string y asígnale tu nombre. Luego imprime el valor de la variable por la consola.
package main

import (
	"fmt"
)

func main() {
	var nombre string
	var edad int
	var pi float32
	nombre = "georgiana"
	edad = 28
	pi = 3.14159

	fmt.Println("El nombre de esta persona es: ", nombre, "\n y la edad es: ", edad, "\n y el numero pi es: ", pi)
}