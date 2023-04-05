//1. Crea un programa que pida al usuario dos números y luego determine cuál es el mayor de los dos.

package main

import (
    "fmt"
)

func main() {
	var num1 int
	var num2 int
	fmt.Print("Ingrese un numero:")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese otro numero:")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Println(num1, "es mayor que el numero: ", num2)
	}else if num1 < num2 {
		fmt.Println(num2, "es mayor que el numero:", num1)
	}else{
		fmt.Println("ambos son iguales")
	}
}