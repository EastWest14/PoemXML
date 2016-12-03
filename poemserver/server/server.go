package server

import (
	"fmt"
	"net/http"
)

const STANDARD_RESPONSE = "Hello, world!"

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, STANDARD_RESPONSE)
}
