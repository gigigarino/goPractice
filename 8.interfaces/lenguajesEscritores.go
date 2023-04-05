/*Crea dos interfaces: Lector y Escritor. La interfaz Lector debe tener un método Leer() string y la interfaz 
Escritor debe tener un método Escribir(string). A continuación, crea structs que representen diferentes 
tipos de lectores y escritores, como Archivo, Buffer y BaseDeDatos. Implementa las interfaces Lector y Escritor 
según corresponda para cada uno de estos structs.*/

package main

import (
    "fmt"
)

// Definición de la interfaz Lector
type Lector interface {
    Leer() string
}

// Definición de la interfaz Escritor
type Escritor interface {
    Escribir(string)
}

// Definición de la estructura Archivo
type Archivo struct {
    nombre string
}

// Implementación del método Leer() de la interfaz Lector para la estructura Archivo
func (a *Archivo) Leer() string {
    // Supongamos que aquí se lee el contenido del archivo y se devuelve como un string
    return fmt.Sprintf("Contenido del archivo %s", a.nombre)
}

// Implementación del método Escribir() de la interfaz Escritor para la estructura Archivo
func (a *Archivo) Escribir(contenido string) {
    // Supongamos que aquí se escribe el contenido en el archivo
    fmt.Printf("Contenido escrito en el archivo %s: %s\n", a.nombre, contenido)
}

// Definición de la estructura Buffer
type Buffer struct {
    contenido string
}

// Implementación del método Leer() de la interfaz Lector para la estructura Buffer
func (b *Buffer) Leer() string {
    return fmt.Sprintf("Contenido del buffer: %s", b.contenido)
}

// Implementación del método Escribir() de la interfaz Escritor para la estructura Buffer
func (b *Buffer) Escribir(contenido string) {
    b.contenido = contenido
    fmt.Printf("Contenido escrito en el buffer: %s\n", contenido)
}

// Definición de la estructura BaseDeDatos
type BaseDeDatos struct {
    nombre string
}

// Implementación del método Leer() de la interfaz Lector para la estructura BaseDeDatos
func (bd *BaseDeDatos) Leer() string {
    // Supongamos que aquí se consulta la base de datos y se devuelve el resultado como un string
    return fmt.Sprintf("Resultado de la consulta a la base de datos %s", bd.nombre)
}

// Implementación del método Escribir() de la interfaz Escritor para la estructura BaseDeDatos
func (bd *BaseDeDatos) Escribir(contenido string) {
    // Supongamos que aquí se escribe el contenido en la base de datos
    fmt.Printf("Contenido escrito en la base de datos %s: %s\n", bd.nombre, contenido)
}

// Función principal que crea instancias de los lectores y escritores y los utiliza
func main() {
    archivo := Archivo{"mi_archivo.txt"}
    buffer := Buffer{""}
    baseDeDatos := BaseDeDatos{"mi_base_de_datos"}

    contenidoArchivo := archivo.Leer()
    fmt.Println(contenidoArchivo)

    buffer.Escribir("Contenido del buffer")
    contenidoBuffer := buffer.Leer()
    fmt.Println(contenidoBuffer)

    baseDeDatos.Escribir("Contenido de la base de datos")
    contenidoBaseDeDatos := baseDeDatos.Leer()
    fmt.Println(contenidoBaseDeDatos)
}