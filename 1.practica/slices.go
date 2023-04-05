package main

import "fmt"

func main (){
	meses := []string{"enero", "diciemnbe", "mayo", "abril", "octubre", "arbil"}
	ObtenerEstaciones(meses)
}

func ObtenerEstaciones(meses []string){
	for _, mes := range meses {
		if mes == "enero" || mes == "febrero" || mes == "marzo"{
			fmt.Println("corresponde a la estacion verano")
		}else if mes == "abril" || mes == "mayo" || mes == "junio" {
			fmt.Println("corresponde a la estacion otoño")
		}else if mes == "julio" || mes == "agosto" || mes == "septiembre"{
			fmt.Println("corresponde a la estacion invierno")
		}else if mes == "octubre" || mes == "noviembre" || mes == "diciembre"{
			fmt.Println("Corresponde a la estacion primaveraaa")
		}else{
			fmt.Println("NO corresponde a ninguna estacion")
		}
	}
}