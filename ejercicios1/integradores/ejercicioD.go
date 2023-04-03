//1. Crea un programa que simule un juego de adivinanza. El programa debe generar un número aleatorio entre 1 y 100, 
//y luego preguntar al usuario que adivine el número. Si el número que ingresó el usuario es mayor que el número aleatorio, 
//el programa debe imprimir "El número es menor". Si el número que ingresó el usuario es menor que el número aleatorio,
// el programa debe imprimir "El número es mayor". Si el número que ingresó el usuario es igual al número aleatorio, 
//el programa debe imprimir "¡Adivinaste!" y terminar el juego.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Inicializamos la semilla del generador de números aleatorios
    rand.Seed(time.Now().UnixNano())

    // Generamos un número aleatorio entre 1 y 100
    numeroAleatorio := rand.Intn(100) + 1

    // Inicializamos una variable para almacenar el número que ingresa el usuario
    var numeroIngresado int

    // Iteramos hasta que el usuario adivine el número
    for {
        // Pedimos al usuario que ingrese un número
        fmt.Print("Ingresa un número entre 1 y 100: ")
        fmt.Scanln(&numeroIngresado)

        // Comparamos el número ingresado por el usuario con el número aleatorio
        if numeroIngresado < numeroAleatorio {
            fmt.Println("El número es mayor")
        } else if numeroIngresado > numeroAleatorio {
            fmt.Println("El número es menor")
        } else {
            fmt.Println("¡Adivinaste!")
            break
        }
    }
}