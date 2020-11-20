package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{rules: make(map[string]map[string]http.HandlerFunc)}
}

func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exists := r.rules[path]
	handler, methodExists := r.rules[path][method]
	return handler, methodExists, exists
}

func (r *Router) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	handler, methodExists, exists := r.FindHandler(request.URL.Path, request.Method)
	if !methodExists {
		writter.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else if !exists {
		writter.WriteHeader(http.StatusNotFound)
		return
	}
	handler(writter, request)
	return
}
