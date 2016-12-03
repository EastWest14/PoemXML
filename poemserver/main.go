package main

import (
	"net/http"
	"poemXML/poemserver/server"
)

func main() {
	http.HandleFunc("/", server.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
