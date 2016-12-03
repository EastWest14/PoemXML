package main

import (
	"net/http"
	"poemXML/poemserver/server"
)

func main() {
	http.HandleFunc("/", server.HandleDefaultRequest)
	http.HandleFunc("/poems", server.HandlePoemsRequest)
	http.ListenAndServe(":8080", nil)
}
