//a) Crea una estructura llamada "persona" que contenga los campos "nombre" y "edad". Luego crea una variable de tipo "persona" 
//y asígnale tus datos. Luego imprime los datos de la persona por la consola.

package main

import (
    "fmt"
)

type Persona struct{
	nombre string
	edad int
    comidafavorita string
}

func main() {
	persona1 := Persona{
        nombre: "Juan",
        edad: 1,
        comidafavorita: "milanesa",
    }

    fmt.Println(persona1.nombre, persona1.comidafavorita)

    persona2 := Persona{
        nombre: "Juan",
        edad: 2,
        comidafavorita: "asado",
    }

    fmt.Println(persona2)

    persona3 := Persona{
        nombre: "Juan",
        edad: 3,
        comidafavorita: "chinchulin",
    }

    fmt.Println(persona3)

    persona4 := Persona{
        nombre: "Juan",
        edad: 4,
        comidafavorita: "pizza",
    }

    fmt.Println(persona4)
}