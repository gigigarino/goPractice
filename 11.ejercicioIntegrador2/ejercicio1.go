package main

/*
Escribe una función en Go que tome un slice de enteros como argumento y devuelva 
la suma de todos los números en el slice.*/

/*Luego, define una interfaz llamada Operacion que tenga un método 
llamado Operar que tome un slice de enteros y devuelva un resultado de algún tipo.*/

/*A continuación, crea dos tipos de estructuras para representar operaciones
 matemáticas: Suma y Promedio. Cada uno debe implementar la interfaz Operacion y tener su propio método Operar.
*/

/*La función principal debe tomar un slice de enteros y una instancia de la 
interfaz Operacion, y luego llamar al método Operar de la instancia de Operacion para obtener
 un resultado. Finalmente, la función principal debe imprimir el resultado obtenido.
 */

 type Operacion interface{
	Operar(slice []int)
 }