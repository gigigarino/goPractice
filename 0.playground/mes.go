/*
realizar una app que contenga una variabke con el numero del mes.
segun el numero, imprimir el mes que corresponda en texto
se te ocurre si se puede resolver de mas de una manera?
*/

package main

import "fmt"
func main(){
	var mes int=5
	mesTexto := ""

	switch mes {
	case 1:
		mesTexto = "enero"
	case 2:
		mesTexto = "febrero"
	case 3:
		mesTexto = "marzo"
	case 4:
		mesTexto = "abril"	
	case 5:
		mesTexto = "mayo"
	case 6:
		mesTexto = "junio"
	case 7:
		mesTexto = "julio"	
	case 8:
		mesTexto = "agosto"
	case 9:
		mesTexto = "septiembre"
	case 10:
		mesTexto = "octubre"
	case 11:
		mesTexto = "noviembre"
	case 12:
		mesTexto = "diciembre"
	default:
		mesTexto = ""	
	}

	fmt.Println(mesTexto)

	/*mesesTexto := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "octubre", "noviembre", "diciembre"}
	for i := 0; i < len(mesesTexto); i++ {
		if i == mes {
			fmt.Println(mesesTexto[i-1])
		}
	}
	*/
}