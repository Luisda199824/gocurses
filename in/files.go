package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content := []byte("Go 1\nGo 2\nGo 3")
	filename := "file.txt"

	writeInFile(filename, content)
}

func writeInFile(filename string, content []byte) {
	datos := ioutil.WriteFile(filename, content, 0755)

	mostrarError(datos)

	fmt.Println("Success")
}

func openFile(path string) {
	datos, err := ioutil.ReadFile(path)

	mostrarError(err)
	values := string(datos)

	fmt.Println(values)
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
