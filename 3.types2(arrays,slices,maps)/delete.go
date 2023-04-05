package main

import (
	"fmt"
)

func main() {
	x := []int{5, 34, 2, 55, 311, 545, 2342, 22, 555555}

	// se unen x[0,2] u x(4,hasta el final de x]
	x = append(x[:2], x[4:]...)

	fmt.Println(x)

}
