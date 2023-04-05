//a) Crea una estructura llamada "persona" que contenga los campos "nombre" y "edad". Luego crea una variable de tipo "persona" 
//y asígnale tus datos. Luego imprime los datos de la persona por la consola.

package main

import (
    "fmt"
)

type Persona struct{
	nombre string
	edad int
}

func main() {
	persona1 := Persona{
        nombre: "Juan",
        edad: 1,
    }

    fmt.Println(persona1)

    persona2 := Persona{
        nombre: "Juan",
        edad: 2,
    }

    fmt.Println(persona2)

    persona3 := Persona{
        nombre: "Juan",
        edad: 3,
    }

    fmt.Println(persona3)

    persona4 := Persona{
        nombre: "Juan",
        edad: 4,
    }

    fmt.Println(persona4)
}