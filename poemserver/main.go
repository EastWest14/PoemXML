package main

import (
	"poemXML/poemserver/handlers"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

func main() {
	poemServer = server.NewServer()

	handlers := handlers.NewHandlersInstance()
	poemServer.SetHandlers(handlers)

	poemStore := poemstore.NewStore()
	handlers.SetPoemStore(poemStore)

	poemServer.Start()
}
