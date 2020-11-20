package main

import (
	"fmt"
)

func main() {

	var coffe Cafe = Cafe{
		nombre: "Expresso",
		precio: 12.5,
		azucar: false,
		leche:  30,
	}

	fmt.Println(coffe)

}

type Cafe struct {
	nombre string
	precio float32
	azucar bool
	leche  int
}
