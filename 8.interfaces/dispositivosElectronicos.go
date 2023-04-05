//Crea una interfaz llamada Dispositivo que tenga métodos como Encender(), 
//Apagar() y Estado() string. Luego, crea structs que representen diferentes dispositivos electrónicos, 
//como Televisor, Computadora y Smartphone. Implementa la interfaz Dispositivo para cada uno de estos structs.

package main

import (
    "fmt"
)

// Definición de la interfaz Dispositivo
type Dispositivo interface {
    Encender()
    Apagar()
    Estado() string
}

// Definición de la estructura Televisor
type Televisor struct {
    encendido bool
    marca     string
}

// Implementación de los métodos de la interfaz Dispositivo para la estructura Televisor
func (t *Televisor) Encender() {
    t.encendido = true
}

func (t *Televisor) Apagar() {
    t.encendido = false
}

func (t *Televisor) Estado() string {
    if t.encendido {
        return fmt.Sprintf("El televisor de marca %s está encendido.", t.marca)
    } else {
        return fmt.Sprintf("El televisor de marca %s está apagado.", t.marca)
    }
}

// Definición de la estructura Computadora
type Computadora struct {
    encendido bool
    marca     string
}

// Implementación de los métodos de la interfaz Dispositivo para la estructura Computadora
func (c *Computadora) Encender() {
    c.encendido = true
}

func (c *Computadora) Apagar() {
    c.encendido = false
}

func (c *Computadora) Estado() string {
    if c.encendido {
        return fmt.Sprintf("La computadora de marca %s está encendida.", c.marca)
    } else {
        return fmt.Sprintf("La computadora de marca %s está apagada.", c.marca)
    }
}

// Definición de la estructura Smartphone
type Smartphone struct {
    encendido bool
    marca     string
}

// Implementación de los métodos de la interfaz Dispositivo para la estructura Smartphone
func (s *Smartphone) Encender() {
    s.encendido = true
}

func (s *Smartphone) Apagar() {
    s.encendido = false
}

func (s *Smartphone) Estado() string {
    if s.encendido {
        return fmt.Sprintf("El smartphone de marca %s está encendido.", s.marca)
    } else {
        return fmt.Sprintf("El smartphone de marca %s está apagado.", s.marca)
    }
}

// Función principal que crea instancias de los dispositivos y los utiliza
func main() {
    tv := Televisor{false, "Samsung"}
    comp := Computadora{false, "Lenovo"}
    phone := Smartphone{false, "Apple"}

    tv.Encender()
    fmt.Println(tv.Estado())

    comp.Encender()
    fmt.Println(comp.Estado())

    phone.Encender()
    fmt.Println(phone.Estado())

    tv.Apagar()
    fmt.Println(tv.Estado())
}