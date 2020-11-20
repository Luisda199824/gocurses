package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerRoot(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "Hello world, root path")
}

func HandlerHome(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "Hello world, from home path")
}

func PostHandlerUser(writter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Fprintf(writter, "error: %v", err)
		return
	}
	fmt.Fprintf(writter, "Payload: %v\n", metadata)
}

func HandlerUser(writter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(writter, "error: %v", err)
		return
	}

	response, errJson := user.ToJSON()

	if errJson != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		return
	}
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	writter.Write(response)
}
