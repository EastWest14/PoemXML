package main

import (
	"net/http"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

func main() {
	setup()
	r := server.CreateAndConfigureRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func setup() {
	poemStore := poemstore.NewStore()
	server.SetPoemStore(poemStore)
}
