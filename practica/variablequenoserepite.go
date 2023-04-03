package main

import (
	"fmt"
)

func main(){
	var a int
	b := 1
	
	
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	a = 3

	c := a +b

	fmt.Printf("c: %T\n", c)
	fmt.Println("el valor de c es: " , c)
}