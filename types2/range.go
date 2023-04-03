package main

import (
	"fmt"
)

func main() {
	x := []int{5, 34, 2, 55, 311}

	fmt.Println("Longitud:", len(x))

	// imprime el slice [valores]
	fmt.Println(x)

	// acceso por index pos
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	fmt.Println("---------------------")

	for i, valorPosSlice := range x {
		fmt.Println("Pos:", i, " Valor:", valorPosSlice)
	}
}
