package handlers

import (
	"fmt"
	"net/http"
	"poemXML/poemserver/poemstore"
)

type Handlers struct {
	poemStore *poemstore.Store
}

func NewHandlersInstance() *Handlers {
	return &Handlers{}
}

func (h *Handlers) SetPoemStore(store *poemstore.Store) {
	h.poemStore = store
}

const STANDARD_RESPONSE = "Hello, world!"

func (h *Handlers) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, STANDARD_RESPONSE)
}

func (h *Handlers) PoemListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of poems")
}
