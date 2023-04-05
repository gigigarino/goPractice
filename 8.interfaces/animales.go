/* Crea una interfaz llamada Animal que tenga métodos como Hablar() string, 
Comer() string y Nombre() string. Luego, crea structs que representen diferentes animales, 
como Perro, Gato y Pajaro. Implementa la interfaz Animalpara cada uno de estos structs.*/

package main

import (
    "fmt"
)

// Definición de la interfaz Animal
type Animal interface {
    Hablar() string
    Comer() string
    Nombre() string
}

// Definición de la estructura Perro
type Perro struct {
    nombre string
}

// Implementación del método Hablar() de la interfaz Animal para la estructura Perro
func (p *Perro) Hablar() string {
    return "Guau"
}

// Implementación del método Comer() de la interfaz Animal para la estructura Perro
func (p *Perro) Comer() string {
    return "Carne"
}

// Implementación del método Nombre() de la interfaz Animal para la estructura Perro
func (p *Perro) Nombre() string {
    return p.nombre
}

// Definición de la estructura Gato
type Gato struct {
    nombre string
}

// Implementación del método Hablar() de la interfaz Animal para la estructura Gato
func (g *Gato) Hablar() string {
    return "Miau"
}

// Implementación del método Comer() de la interfaz Animal para la estructura Gato
func (g *Gato) Comer() string {
    return "Pescado"
}

// Implementación del método Nombre() de la interfaz Animal para la estructura Gato
func (g *Gato) Nombre() string {
    return g.nombre
}

// Definición de la estructura Pajaro
type Pajaro struct {
    nombre string
}

// Implementación del método Hablar() de la interfaz Animal para la estructura Pajaro
func (p *Pajaro) Hablar() string {
    return "Pio pio"
}

// Implementación del método Comer() de la interfaz Animal para la estructura Pajaro
func (p *Pajaro) Comer() string {
    return "Semillas"
}

// Implementación del método Nombre() de la interfaz Animal para la estructura Pajaro
func (p *Pajaro) Nombre() string {
    return p.nombre
}

// Función principal que crea instancias de los animales y los utiliza
func main() {
    perro := Perro{"Fido"}
    gato := Gato{"Misi"}
    pajaro := Pajaro{"Piolín"}

    fmt.Printf("El %s dice: %s\n", perro.Nombre(), perro.Hablar())
    fmt.Printf("El %s come: %s\n", perro.Nombre(), perro.Comer())

    fmt.Printf("La %s dice: %s\n", gato.Nombre(), gato.Hablar())
    fmt.Printf("La %s come: %s\n", gato.Nombre(), gato.Comer())

    fmt.Printf("El %s dice: %s\n", pajaro.Nombre(), pajaro.Hablar())
    fmt.Printf("El %s come: %s\n", pajaro.Nombre(), pajaro.Comer())
}