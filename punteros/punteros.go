package main

import "fmt"

func main() {
    // Declarar y asignar una variable
    numero := 42

    // Obtener la dirección de memoria de la variable
    puntero := &numero

    // Imprimir la dirección de memoria
    fmt.Printf("Dirección de memoria de numero: %p\n", puntero)

    // Acceder al valor almacenado en la dirección de memoria apuntada por el puntero
    valor := *puntero
    fmt.Printf("Valor apuntado por puntero: %d\n", valor)

    // Modificar el valor almacenado en la dirección de memoria apuntada por el puntero
    *puntero = 99
    fmt.Printf("Nuevo valor de numero: %d\n", numero)
}
