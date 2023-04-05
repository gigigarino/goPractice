//1. Crea un programa que pregunte al usuario su edad y luego determine si es mayor de edad o no.

package main

import (
    "fmt"
)

func main() {
	var edad int
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	if edad >= 18 {
		fmt.Println("usted tiene la edad de: ", edad, "años, puede ingresar")
	} else {
		fmt.Println("usted tiene la edad de: ", edad, "años, NO PUEDE INGRESAR")
	}
}