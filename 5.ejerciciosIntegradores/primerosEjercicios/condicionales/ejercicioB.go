//b) Crea una variable "edad" de tipo int y asígnale tu edad. Luego crea un condicional que imprima 
//"Eres mayor de edad" si la edad es mayor o igual a 18 años, y 
//"Eres menor de edad" si la edad es menor que 18 años. 

package main

import (
    "fmt"
)

func main() {
	edad := 28

	if edad >= 18 {
		fmt.Println("sos mayor de edad, entra a revolear el toto en el boliche")
	}else{
		fmt.Println("sos menor de edad, anda a tu casa")
	}
}