package main

import (
	"poemXML/poemserver/handlers"
	"poemXML/poemserver/index"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

const INDEX_PATH = "./index"

func main() {
	poemServer = server.NewServer()

	handlers := handlers.NewHandlersInstance()
	poemServer.SetHandlers(handlers)

	theIndex := index.New(INDEX_PATH)

	poemStore := poemstore.NewStore(theIndex)
	handlers.SetPoemStore(poemStore)

	poemServer.Start()
}
