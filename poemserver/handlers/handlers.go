package handlers

import (
	"fmt"
	"net/http"
)

type Handlers struct {
}

func NewHandlersInstance() *Handlers {
	return &Handlers{}
}

const STANDARD_RESPONSE = "Hello, world!"

func (h *Handlers) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, STANDARD_RESPONSE)
}

func (h *Handlers) PoemListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of poems")
}
