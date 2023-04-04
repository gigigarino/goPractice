package main

import (
	"fmt"
)

func main (){
	foo()
	bar("hello word")
	fmt.Println(woo("woo()"))
	fmt.Println(moo("Bart", "Simpson"))
	fmt.Println(moo("Homero", "Simpson"))
}

func foo(){
	fmt.Println("hola desde foo")
}

func bar(s string){
	fmt.Println(s)
}
func woo(w string) string {
	w2 := "Hola desde " + w + "!"

	return w2
}

func moo(nombre string, apellido string) (string, bool) {

	acreditado := false
	posicion := "x"

	if nombre == "Homero" {
		if apellido == "Simpson" {
			acreditado = true
			posicion = "sector 7g"
		}
	}

	return posicion, acreditado
}
