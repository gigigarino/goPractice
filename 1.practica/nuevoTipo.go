package main

import (
	"fmt"
)

var a int

type miTipo int //se crea un nuevo tipo

var p miTipo

func main() {

	a = 11
	fmt.Printf("El tipo de a es: %T\n", a)
	fmt.Println("El valor de a es:", a)

	p = 22
	fmt.Printf("El tipo de p es: %T\n", p)
	fmt.Println("El valor de p es: ", p)

	//a = p //esto dará error pq son de distintos tipos!
	p = 33
	a = int(p) // funciona pq p es convertida en int
	fmt.Printf("El tipo de a es: %T\n", a)
	fmt.Println("El valor de a es: ", a) // en otros lenguajes esto se llama casting, pero en go es converting

	a = 44
	p = miTipo(a)
	fmt.Printf("El tipo de p es: %T\n", p)
	fmt.Println("El valor de p es: ", p)

}