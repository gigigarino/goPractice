/*
ver la cantidad de letras que tiene una palabra entera
*/

package main

import "fmt"

func main(){
	palabra := "fisioterapia"
	var cant int = len(palabra)

	fmt.Printf("la palabra %s tiene %d letras ", palabra, cant)
	//es la parte mas facil de decir
	//fmt.Print("la palabra", palabra, "tiene", cant, "letras")

	//imprimir cada una de las letras con un for

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i]))
	}
}

