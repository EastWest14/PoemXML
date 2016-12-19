package handlers

import (
	"fmt"
	"github.com/EastWest14/errorcode"
	"github.com/EastWest14/gAssert"
	"net/http"
	"poemXML/poemserver/poemlist"
	"poemXML/poemserver/poemstore"
)

const (
	DATA_ACCESS_FAILED_USER_MESSAGE = "Problem accessing poem data"
	UNKNOWN_ERROR_USER_MESSAGE      = "Unknown server error"
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

	listOfAllPoems, err := h.poemStore.GetAllPoems()
	if err == nil {
		fmt.Fprint(w, listOfAllPoems.String())
		return
	}

	switch err.Type() {
	case poemstore.INDEX_UNAVAILABLE_ERROR:
		http.Error(w, DATA_ACCESS_FAILED_USER_MESSAGE, http.StatusInternalServerError)
	default:
		http.Error(w, UNKNOWN_ERROR_USER_MESSAGE, http.StatusInternalServerError)
	}
}

type PoemStoreT interface {
	GetAllPoems() (list *poemlist.PoemList, err *errorcode.Errorcode)
}
