/*
Creación de errores

La biblioteca estándar ofrece dos funciones incorporadas para crear errores: errors.New y fmt.Errorf..

errors.New tiene un solo argumento: un mensaje de error con una cadena que puede personalizar 
para avisarles a sus usuarios cuál fue el problema.*/


package main

import (
	"fmt"
	"errors"
)

func main() {
	err := errors.New("bernacles")
	fmt.Println("sammy says:", err)
}