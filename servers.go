package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	for _, url := range servers {
		reviewServer(url)
	}

	tiempoEjecucion := time.Since(inicio)

	fmt.Printf("Tiempo de ejecución: %s", tiempoEjecucion)
}

func reviewServer(url string) bool {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Servidor " + url + " no disponible")
		return true
	}
	fmt.Println("Servidor " + url + " disponible")
	return false
}
