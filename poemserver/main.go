package main

import (
	"net/http"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

func main() {
	setup()
	poemServer = server.NewServer()
	r := poemServer.CreateAndConfigureRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func setup() {
	poemStore := poemstore.NewStore()
	poemServer.SetPoemStore(poemStore)
}
