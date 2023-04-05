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
    "sort"
)

type Ordenable interface {
    MenorQue(otro Ordenable) bool
}

type Persona struct {
    nombre string
    edad   int
}

func (p Persona) MenorQue(otro Ordenable) bool {
    return p.edad < otro.(Persona).edad
}

type Producto struct {
    nombre string
    precio float64
}

func (p Producto) MenorQue(otro Ordenable) bool {
    return p.precio < otro.(Producto).precio
}

func Ordenar(elementos []Ordenable) {
    sort.Slice(elementos, func(i, j int) bool {
        return elementos[i].MenorQue(elementos[j])
    })
}

func main() {
    personas := []Persona{
        {"Juan", 30},
        {"Maria", 25},
        {"Pedro", 40},
    }
    Ordenar([]Ordenable{})
    fmt.Println(personas)

    productos := []Producto{
        {"Libro", 20.0},
        {"Lapiz", 1.0},
        {"Papel", 5.0},
    }
    Ordenar([]Ordenable{})
    fmt.Println(productos)
}

