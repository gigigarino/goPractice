//Crea una interfaz llamada Forma que tenga dos métodos: Area() y Perimetro(). 
//Luego, crea structs que representen diferentes formas geométricas, como Cuadrado, Rectangulo y Circulo. 
//Implementa la interfaz Forma para cada uno de estos structs ne lenguaje go


package main

import(
	"fmt"
	"math"
)

type Forma interface {
	area() float64
	perimetro() float64
}

type cuadrado struct {
	lado float64
}

// Implementamos la interfaz Forma para el struct Cuadrado
func (c cuadrado) Area() float64 {
    return c.lado * c.lado
}

func (c cuadrado) Perimetro() float64 {
    return 4 * c.lado
}

type rectangulo struct {
	base float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func (r rectangulo) perimetro() float64{ 
	return 2*r.base * 2*r.altura
}

//circulo
type circulo struct {
	radio float64
}

func (c circulo)perimetro() float64{
	return 2* math.Pi * c.radio
}

func (c circulo) area()float64{
	return math.Pi * c.radio * c.radio
}

func formasgeometricas() {
	cuadrado := cuadrado{lado: 5}
	rectangulo := rectangulo{base:10, altura:8}
	circulo := circulo{radio:3}

	fmt.Println("area del cuadrado: ", cuadrado.Area())
	fmt.Println("area del rectangulo: ", rectangulo.area())
	fmt.Println("area del circulo: ", circulo.area())
	fmt.Println("perimetro del cuadrado: ", cuadrado.Perimetro())
	fmt.Println("perimetro del rectangulo: ",rectangulo.perimetro())
	fmt.Println("area del circulo: ", circulo.perimetro())
}