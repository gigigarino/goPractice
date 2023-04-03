//1. Crea un programa que pida al usuario un número y luego determine si es par o impar.

package main

import (
    "fmt"
)

func main() {
	var numero int
	fmt.Print("ingresa un numero: ")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Println("el numero es par")
	}else{
		fmt.Println("el numero es impar")
	}


}