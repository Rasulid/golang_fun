package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, m)
	//TODO implement me
	//panic("implement me")
}

func (m msg) ServerHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}

func main() {
	msgHandler := msg("Server HTTP Listen Address")
	http.ListenAndServe("localhost:8181", msgHandler)
}
