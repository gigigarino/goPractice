package main

import (
	"fmt"
)

type Persona struct {
	nombre string
	edad int
}

func aumentarEdad (p *Persona){
	p.edad++
}

func main(){
	persona := Persona{nombre: "georgiana", edad: 28}
	fmt.Printf("antes de aumentar la edad: %s, %d\n", persona.nombre, persona.edad)

	aumentarEdad(&persona)
	fmt.Printf("despues de aumentar la edad: %s, %d\n", persona.nombre, persona.edad)
}