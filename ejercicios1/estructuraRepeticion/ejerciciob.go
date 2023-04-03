//b) Crea un slice de tipo int con los números del 1 al 5. 
//Luego crea un loop que imprima los elementos del slice por la consola.

package main

import "fmt"

func main() {
    numeros := []int{1, 2, 3, 4, 5}

    for _, numero := range numeros {
        fmt.Println(numero)
    }
}