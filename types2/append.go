package main

import (
	"fmt"
)

func main() {
	unSliceDeInts := []int{5, 34, 2, 55, 311}
	fmt.Println("unSliceDeInts antes de append:", unSliceDeInts)

	unSliceDeInts = append(unSliceDeInts, 44, 11, 55, 1231)
	fmt.Println("unSliceDeInts despues de append:", unSliceDeInts) // se adjuntan valores al slice

	otroSliceDeInts := []int{88, 33453, 11, 657, 44444}

	unSliceDeInts = append(unSliceDeInts, otroSliceDeInts...)

	fmt.Println("Se adjunto todo el otroSliceDeInts en unSliceDeInts", unSliceDeInts)
}
