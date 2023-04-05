/*La función fmt.Errorf le permite crear un mensaje de error de forma dinámica.
Su primer argumento es una cadena que contiene su mensaje de error con valores de marcadores de posición,
como %s para cadenas y %d para enteros. fmt.Errorf interpola los argumentos que siguen esta cadena
de formato en esos marcadores de posición en orden:*/

package main

import (
	"fmt"
	"time"
)

func main(){
	err := fmt.Errorf("error ocurrido en: %v", time.Now())
	fmt.Println("ocurrio un error", err)
}