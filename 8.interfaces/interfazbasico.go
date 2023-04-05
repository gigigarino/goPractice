package main

import (
	"fmt"
)

//interfaz
type Hablante interface {
	Hablar () string
}

//estructura
type Persona struct {
	nombre string
}

func (p Persona) Hablar () string {
	return "hello, mi nombre es " + p.nombre
}

//estructura
type Gato struct {
	nombre string
}


func (g Gato) Hablar() string {
	return "Miau, soy " + g.nombre
}

func Saludar(h Hablante) {
	fmt.Println(h.Hablar())
}

func main() {
	persona := Persona{nombre: "Juan"}
	gato := Gato{nombre: "Whiskers"}

	Saludar(persona)
	Saludar(gato)
}