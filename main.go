package main // Nombre del paquete

import "fmt" // Importar todas las librerias necesarias

func main() {
	var mensaje string = "Hola mundo" // var NAME_VAR type_var = value
	mensaje2 := "Hola mundo" // Uso de :=
	fmt.Println(mensaje)
	fmt.Println(mensaje2)
	a := 10.
	var b float64 = 2.
	fmt.Println(a/b)
	var c int = 10
	d := 3
	fmt.Println(c/d)
	x := true
	y := false
	fmt.Println(!x)
	fmt.Println(y)
	fmt.Println(x && y)
}

func privada() {
	fmt.Println("Ejecuta solo dentro del paquete main por ser privada")
}

func Publica() {
	fmt.Println("Ejecuta global")
}

func imprimirHola() {
	defer fmt.Println("Mundo") // Ejecuta hasta el final de la función
	fmt.Println("Hola")
}