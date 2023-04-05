//b) Crea una estructura llamada "producto" que contenga los campos "nombre", "precio" y "cantidad". 
//Luego crea una variable de tipo "producto" y asígnale los datos de un producto cualquiera. 
//Luego imprime los datos del producto por la consola

package main

import (
    "fmt"
)

type Producto struct {
	nombre string
	precio float64
	cantidad int
}

func main() {
	producto := Producto{
        nombre: "limpaividrio",
        precio: 0.1,
        cantidad: 1,
    }

    fmt.Println("Producto: ", producto)
}