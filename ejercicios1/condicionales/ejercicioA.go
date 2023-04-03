// Crea una variable "numero" de tipo int y asígnale un número cualquiera.
// Luego crea un condicional que imprima "Es un número positivo" si el número es mayor que cero,
// "Es un número negativo" si el número es menor que cero, y "El número es cero" si el número es igual a cero.
package main

import (
	"fmt"
)

func main() {
	var numero int
	numero = 23

	if numero > 0 {
		fmt.Println("el numero es positivo")
	} else if numero < 0 {
		fmt.Println("el numero es negativo")
	} else {
		fmt.Println("el numero es igual a cero")

	}
}
