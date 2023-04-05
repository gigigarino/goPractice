package main

import (
	"fmt"
)

func main() {
	var a int // otra forma de declarar una variable, asigna el tipo y valor por defecto del tipo
	b := 12   // short operator, asigna valor y tipo

	fmt.Printf("a: %T\n", a) // imprime el tipo de a
	fmt.Printf("b: %T\n", b) // imprime el tipo de b

	a = 3

	c := a + b // se asigna el valor de la suma a c y al mismo se declara c, como el resultado es de tipo int, c tambien es de tipo int

	fmt.Println("El valor de c es:", c)
	fmt.Printf("c: %T\n", c)

    //c := 23 // INCORRECTO,como se ya fue declarada NO se debe usar el operardor := 
    c = 23 // correcto, de esta forma se cambiar el valor de c
    fmt.Println("El valor de c es:", c)
}
