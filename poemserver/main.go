package main

import (
	//"poemXML/poemserver/poemstore"
	"poemXML/poemserver/handlers"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

func main() {
	poemServer = server.NewServer()
	handlers := handlers.NewHandlersInstance()
	poemServer.SetHandlers(handlers)
	poemServer.Start()
}
