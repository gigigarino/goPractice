package main

import (
	"fmt"
)

func main(){
	a := 45
	fmt.Println("var a (valor de la variable a): ", a)
	fmt.Println("var &a (la dire de memoria donde se almaena el valor de a)", &a)

	fmt.Println("--------------------------------")
	fmt.Printf("a es de tipo: %T\n", a)
	fmt.Printf("&a es de tipo: %T\n", &a)
	fmt.Println("--------------------------------")


	b := &a
	fmt.Println("b almacena la direccion de b:", b)
	fmt.Println("El valor que apunta el puntero:", *b) // * devuelve el valor almacenado
	fmt.Println("El valor que apunta el puntero:", *&a)

	fmt.Println("--------------------")


	*b = a
	fmt.Println("b apunta a la direccion de memoria de a:", b) // a y *b estan alcenandos en la misma posición de memoria
	fmt.Printf("b es de tipo: %T\n", b)
	fmt.Println("--------------------")

	*b = 555
	fmt.Println("Al cambiar el valor *b a 555, tambien cambia a:", a) // a y *b estan alcenandos en la misma posición de memoria
}