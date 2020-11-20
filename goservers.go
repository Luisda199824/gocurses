package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	channel := make(chan string)
	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	counter := 0

	for {
		if counter > 10 {
			break
		}

		for _, url := range servers {
			go reviewServer(url, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		counter++
	}

	tiempoEjecucion := time.Since(inicio)

	fmt.Printf("Tiempo de ejecución: %s", tiempoEjecucion)
}

func reviewServer(url string, channel chan string) bool {
	_, err := http.Get(url)
	if err != nil {
		channel <- "Servidor " + url + " no disponible"
		return true
	}
	channel <- "Servidor " + url + " disponible"
	return false
}
