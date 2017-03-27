package main

import (
	"flag"
	"fmt"
	"github.com/EastWest14/gAssert"
	"poemXML/poemserver/database"
	"poemXML/poemserver/handlers"
	"poemXML/poemserver/poemstore"
	"poemXML/poemserver/server"
)

var poemServer *server.Server

func main() {
	//fmt.Printf("Index path: %s\n", indexPath) //Deprecated
	indexPath := ""
	poemServer := Setup(indexPath)

	fmt.Println("Starting server...")
	poemServer.Start()
}

func ExtractDBConfig() (indexPath string, dbConfig *database.DBConfig) {
	flag.Parse()
	numArgs := flag.NArg()
	gAssert.AssertHard(numArgs > 0, "Number of command line arguments is 0 - can't extract index path.")
	fmt.Printf("Number of arguments: %d\n", flag.NArg())
	gAssert.AssertHard(numArgs >= 4, "Number of command line arguments is < 4 - missing DB parameters.")

	host := flag.Arg(1)
	dbUser := flag.Arg(2)
	dbName := flag.Arg(3)
	dbPassword := flag.Arg(4)

	dbConfig = database.NewDBConfig(host, dbUser, dbName, dbPassword)

	return flag.Arg(0), dbConfig
}

func Setup(indexPath string) *server.Server {
	indexPath, dbConfig := ExtractDBConfig()

	db, err := database.NewConnectedDB(dbConfig)
	if err != nil {
		fmt.Println("Error launching DB.")
		panic(err.Error())
	}

	poemServer = server.NewServer()

	handlers := handlers.NewHandlersInstance()
	poemServer.SetHandlers(handlers)

	poemStore := poemstore.NewStore(nil, db)
	handlers.SetPoemStore(poemStore)
	err = poemStore.Check()

	gAssert.Assert(err == nil, fmt.Sprintf("Poem store couldn't be loaded: %v", err))

	return poemServer
}
