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
	indexPath := ExtractEnvVariables()
	fmt.Printf("Index path: %s\n", indexPath)

	poemServer := Setup(indexPath)

	fmt.Println("Starting server...")
	poemServer.Start()
}

func ExtractEnvVariables() string {
	flag.Parse()
	numArgs := flag.NArg()
	gAssert.AssertHard(numArgs > 0, "Number of command line arguments is 0 - can't extract index path.")
	fmt.Printf("Number of arguments: %d\n", flag.NArg())
	gAssert.AssertHard(numArgs >= 4, "Number of command line arguments is < 4 - missing DB parameters.")

	fmt.Println(flag.Arg(1))
	fmt.Println(flag.Arg(2))
	fmt.Println(flag.Arg(3))
	fmt.Println(flag.Arg(4))

	return flag.Arg(0)
}

func Setup(indexPath string) *server.Server {
	poemServer = server.NewServer()

	handlers := handlers.NewHandlersInstance()
	poemServer.SetHandlers(handlers)

	theIndex := index.New(indexPath)

	poemStore := poemstore.NewStore(theIndex)
	handlers.SetPoemStore(poemStore)
	err := poemStore.Check()
	gAssert.Assert(err == nil, fmt.Sprintf("Poem store couldn't be loaded: %v", err))

	return poemServer
}
