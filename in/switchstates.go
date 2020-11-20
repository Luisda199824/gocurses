package main

import (
	"fmt"
	"time"
)

func main() {
	tiempo := time.Now()
	diaHoy := tiempo.Weekday()

	switch diaHoy {
	case 0:
		fmt.Printf("Domingo")
	case 1:
		fmt.Printf("Lunes")
	case 2:
		fmt.Printf("Martes")
	case 3:
		fmt.Printf("Miércoles")
	case 4:
		fmt.Printf("Jueves")
	case 5:
		fmt.Printf("Viernes")
	case 6:
		fmt.Printf("Sábado")
	}
}
