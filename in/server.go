package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", saludo)
	http.ListenAndServe(":8000", nil)
}

func saludo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola!")
}
