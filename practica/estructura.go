//declarar una persona con nombre y apellido 

//paqutes no es camel case y es una sola palabra y solo minuscula
package main

import (
	"fmt"
)


type Persona struct {
	nombre string
	apellido string 
	edad int
}

type Empleado struct {
	puesto string
	legajo int
	departamento string
	obraSocial string
	//meter caracteristica persona como estructura embebida
	//Persona 
	datosPersonales Persona
	//de esta forma
	//Persona Persona 
}


func main() {
 fmt.Println("helo word")

 p1 := Persona{
    nombre: "fulana de ",
    apellido: "tal",
	edad: 15,
}

fmt.Println(p1)
fmt.Println(cambiarNombre(p1))
x := cambiarNombre(p1)
fmt.Println(x)

emp1 := Empleado{
}
fmt.Println(emp1)




//estructura basica de una funcion
func cambiarNombre(p Persona)Persona {
	p.nombre = "juan"
	return p
}

//funcion
//func (p persona) nombreMetodo ( ...) () {}

//metodo
//func nombreMetodo ( ...) () {}