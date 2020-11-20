package main

import (
	"fmt"
	"strings"

	"github.com/luisda199824/mycalculator"
)

func main() {
	operacion := mycalculator.ReadEntry()

	operador := mycalculator.DetectOperador(operacion)
	valores := strings.Split(operacion, operador)

	// Cast valores from text to number
	operador1 := mycalculator.StrToInt(valores[0])
	operador2 := mycalculator.StrToInt(valores[1])

	calc := mycalculator.Calculator{operador, operador1, operador2}

	fmt.Println("Resultado: ", calc.operar(
		calc.operador,
		calc.operador1,
		calc.operador2,
	))
}
