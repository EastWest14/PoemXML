package main

import (
	"net/http"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

func main() {
	setup()
	//http.HandleFunc("/poems", server.HandlePoemsRequest)
	http.HandleFunc("/", server.HandleDefaultRequest)
	http.ListenAndServe(":8080", nil)
}

func setup() {
	poemStore := poemstore.NewStore()
	server.SetPoemStore(poemStore)
}
