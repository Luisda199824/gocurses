package main

import "fmt"

func main() {
	/* var array [2]int
	array[0] = 1
	array[1] = 2 */

	// array := [2]int{1, 2}

	var frutas [2][2]string
	frutas[0][0] = "manzana"
	frutas[0][1] = "golden"
	frutas[1][0] = "pera"
	frutas[1][1] = "queso"

	fmt.Println(frutas)
}
