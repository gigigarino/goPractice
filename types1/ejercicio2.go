//* Declarar tres variables globales un entero un string y un booleano.

//* Imprimir sus tipos y sus valores por defecto (zeros).

package main

import (
	"fmt"
)

//declaro 3 variables
var a int
var b string
var c bool

func main(){
	fmt.Printf("%T, %T, %T\n",a,b,c)
	fmt.Printf("estos son los 'zeros',los valores por defecto, que asigna el sistema: %v, %v, %v\n\n", a,b,c)


	fmt.Println("zero int",a)
    fmt.Println("zero sting",b)
    fmt.Println("zero bool",c)
}