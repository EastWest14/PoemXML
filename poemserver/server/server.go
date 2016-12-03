package server

import (
	"fmt"
	"net/http"
)

const STANDARD_RESPONSE = "Hello, world!"

func HandleDefaultRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, STANDARD_RESPONSE)
}

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
