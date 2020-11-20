package main

import (
	"fmt"
)

func main() {
	// param1 := 10.0
	// param2 := 0.0
	// fmt.Println(div(param1, param2))

	// name := "Luis"
	// fmt.Println(hi(name))

	// fmt.Println(ganado(1))

	alumnos("Luis", "David", "Luisa", "Andres")
}

func alumnos(alumno ...string) {
	for _, valor := range alumno {
		fmt.Println(valor)
	}
}

func ganado(number int) (int, string) {

	vacas := func() int {
		return number * 10
	}

	return vacas(), "vacas"
}

func div(param1 float64, param2 float64) float64 {
	if param2 == 0 {
		return 0
	}
	return param1 / param2
}

func hi(name string) (string, string) {
	return "Hi ", name
}
