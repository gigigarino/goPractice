//Crea una interfaz llamada Dispositivo que tenga métodos como Encender(), 
//Apagar() y Estado() string. Luego, crea structs que representen diferentes dispositivos electrónicos, 
//como Televisor, Computadora y Smartphone. Implementa la interfaz Dispositivo para cada uno de estos structs.

package main

import (
	"fmt"
)

type Dispositivo interface {
	Enceder() 
	Apagar() 
	Estado() string
}

//dispositivos electronicos
type Televisor struct {
	encendido bool
}

// funciones de el dispositivo
func (t *Televisor) Enceder(){
	t.encendido = true
}

func (t *Televisor) Apagar(){
	t.encendido = false
}

func (t *Televisor) Estado()string {
	if t.encendido {
		return "encendido"
	}else{
		return "apagado"
	}
}

//"--------------------------------"
type Computadora struct {
	encendido bool
}

func (c *Computadora) Enceder(){
	c.encendido = true
}

func (c *Computadora) Apagar(){
	c.encendido = false
}

func 
//"--------------------------------"

type Smartphone struct {}