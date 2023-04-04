package main

import (
	"fmt"
)

type Incrementador interface {
	Incrementar()
}

type Contador struct {
	Valor int
}

func (c *Contador) Incrementar() {
	c.Valor++
}

func IncrementarDosVeces(i Incrementador) {
	i.Incrementar()
	i.Incrementar()
}

func main() {
	contador := &Contador{Valor: 0}

	fmt.Println("Valor inicial del contador:", contador.Valor)

	IncrementarDosVeces(contador)

	fmt.Println("Valor del contador después de incrementar dos veces:", contador.Valor)
}
