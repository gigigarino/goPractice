//Crea una interfaz llamada Forma que tenga dos métodos: Area() y Perimetro(). 
//Luego, crea structs que representen diferentes formas geométricas, como Cuadrado, Rectangulo y Circulo. 
//Implementa la interfaz Forma para cada uno de estos structs ne lenguaje go


package main

import(
	"fmt"
	"math"
)

type Forma interface {
	area() float64
	permitro() float64
}

type cuadrado struct {
	lado float64
}

func permitro (c cuadrado) float64 {
	return c.lado * c.lado
}

func area (c cuadrado) float64 {
	return 4 * c.lado
}

type rectangulo structu