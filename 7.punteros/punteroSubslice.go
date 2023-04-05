package main

import (
	"fmt"
)

func main (){
	//creamos un slice comun
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice original", numeros)

	//creamos un subslice
	subslice := numeros[3:9]
	fmt.Println("subslice", subslice)

	//modificamos el elemento del subslice
	subslice[0] = 99
	subslice[5] = 100

	fmt.Println("subslice modificado", subslice)
	fmt.Println("slice original modificado", numeros)
}