package main

import (
	"fmt"
)

var a int

type miTipo int //se crea un nuevo tipo

var p miTipo

func main() {

	a = 51

	fmt.Println(a)
	fmt.Printf("%T\n", a)


	p = 44
	fmt.Println(p)
	fmt.Printf("%T\n", p) // la salida es main.miTipo porque main es el paquete contiene el tipo miTipo

	// a = p //esto dará error pq son de distintos tipos!
}
