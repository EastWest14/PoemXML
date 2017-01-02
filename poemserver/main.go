package main

import (
	"flag"
	"fmt"
	"github.com/EastWest14/gAssert"
	"poemXML/poemserver/handlers"
	"poemXML/poemserver/index"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

func main() {
	indexPath := ExtractIndexPath()
	fmt.Printf("Index path: %s\n", indexPath)

	poemServer := Setup(indexPath)

	poemServer.Start()
}

func ExtractIndexPath() string {
	flag.Parse()
	numArgs := flag.NArg()
	gAssert.AssertHard(numArgs > 0, "Number of command line arguments is 0 - can't extract index path.")
	fmt.Printf("Number of arguments: %d\n", flag.NArg())

	return flag.Arg(0)
}

func Setup(indexPath string) *server.Server {
	poemServer = server.NewServer()

	handlers := handlers.NewHandlersInstance()
	poemServer.SetHandlers(handlers)

	theIndex := index.New(indexPath)

	poemStore := poemstore.NewStore(theIndex)
	handlers.SetPoemStore(poemStore)

	return poemServer
}
