package main

import (
	"fmt"
	"io"
	"net/http"
)

type writterWeb struct{}

func (writterWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("http://google.com")
	if err == nil {
		fmt.Println(err)
	}

	e := writterWeb{}
	io.Copy(e, res.Body)
}
