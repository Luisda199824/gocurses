package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Nombre   string `json:"nombre,omitempty"`
	Email    string `json:"email,omitempty"`
	Telefono string `json:"telefono,omitempty"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
