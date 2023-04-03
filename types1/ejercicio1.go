//* Usar el short declarator operator para declarar un entero, un string y un booleano.
//* Imprimirlos por pantalla en la misma sentencia.
//* Imprimir cada variable por separado.

package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "Teamcubation!"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}