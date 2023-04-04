//Crea una interfaz llamada Ordenable con un único método: 
//MenorQue(otro Ordenable) bool. Luego, crea structs que 
//representen diferentes tipos de objetos que puedan ser 
//comparados, como Persona (comparando edades) y 
//Producto (comparando precios). 
//Implementa la interfaz Ordenable para cada uno de estos structs. 
//Finalmente, escribe una función genérica Ordenar que acepte un slice de elementos 
//Ordenable y los ordene utilizando el método MenorQue.

package main

import (
	"fmt"
)

type Ordenable interface {
	menorQue(Ordenable) bool
	obtenerValor(Ordenable) int
}

type Persona struct {
	edad int
}

type Producto struct {
	precio int
}

func(p Persona) obtenerValor() int {
	return p.edad
}


func(p Producto) obtenerValor() int {
	return p.precio
}

