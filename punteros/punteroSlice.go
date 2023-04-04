package main

import (
	"fmt"
)

func modificarSlices (s []int){
	s[0] = 100
}

func main (){
	numeros := []int {1,2,3}
	fmt.Println("antes demodificar", numeros)

	modificarSlices(numeros)
	fmt.Println("Despues de moficiar", numeros)
}