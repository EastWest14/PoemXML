package handlers

import (
	"fmt"
	"github.com/EastWest14/gAssert"
	"net/http"
	"poemXML/poemserver/poemlist"
)

type Handlers struct {
	poemStore PoemStoreT
}

func NewHandlersInstance() *Handlers {
	return &Handlers{}
}

func (h *Handlers) SetPoemStore(store PoemStoreT) {
	h.poemStore = store
}

const STANDARD_RESPONSE = "Hello, world!"

func (h *Handlers) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, STANDARD_RESPONSE)
}

func (h *Handlers) PoemListHandler(w http.ResponseWriter, r *http.Request) {
	gAssert.AssertHard(h != nil, "Handlers structure is nil")
	gAssert.AssertHard(h.poemStore != nil, "PoemStore is nil")

	listOfAllPoems := h.poemStore.GetAllPoems()
	fmt.Fprint(w, listOfAllPoems.String())
}

type PoemStoreT interface {
	GetAllPoems() *poemlist.PoemList
}
