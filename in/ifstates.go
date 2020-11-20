package main

import "fmt"

func main() {
	puntos := -9

	if puntos == 10 {
		fmt.Println("Casi ganas")
	} else if puntos > 10 {
		fmt.Println("Ganaste!")
	} else if puntos < 10 && puntos > 0 {
		fmt.Println("Perdiste")
	} else {
		fmt.Println("Impaktado")
	}
}
