package server

import (
	"fmt"
	"net/http"
	"net/url"
	"poemXML/poemserver/poemstore"
	"strings"
)

const STANDARD_RESPONSE = "Hello, world!"

var poemStore *poemstore.Store

func SetPoemStore(pStore *poemstore.Store) {
	poemStore = pStore
}

func HandleDefaultRequest(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	fmt.Println(urlPath)

	fmt.Fprint(w, STANDARD_RESPONSE)
}

func firstPartOfPath(u *url.URL) string {
	urlPath := u.Path
	pathParts := strings.Split(urlPath, "/")

	if len(pathParts) > 1 {
		return pathParts[1]
	}

	return ""
}

/*
func HandlePoemsRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "List of poems")
	case "POST":
		http.Error(w, "Invalid method - use GET!", http.StatusMethodNotAllowed)
	default:
		panic("Invalid method")
	}
}
*/
