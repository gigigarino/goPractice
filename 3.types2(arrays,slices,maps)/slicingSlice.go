package main

import (
	"fmt"
)

func main() {
	x := []int{5, 34, 2, 55, 311}

	fmt.Println("Longitud:", len(x))

	// muestra el slice [valores]
	fmt.Println("Slice completo:", x)

	// slicing a slice, cortando el pedazo
	// muestra pos 0
	fmt.Println("\nPos x[0]:", x[0])

	//// muestra el slice [valores]
	fmt.Println("\nSlice completo x[:]:", x[:])

	// se incluye desde la primera hasta la penultima posición nombrada
	fmt.Println("\nSlice x[1:3]", x[1:3])

	fmt.Println("\nCon range:")
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("\nCon for de 3 valores:")
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println("\nO sea, ambos de arriba son equivalentes.")
}