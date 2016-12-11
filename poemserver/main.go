package main

import (
	//"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

func main() {
	setup()
	poemServer = server.NewServer()
	poemServer.Start()
}

func setup() {
	//poemStore := poemstore.NewStore()
	//poemServer.SetPoemStore(poemStore)
}
