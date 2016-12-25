package handlers

import (
	"fmt"
	"github.com/EastWest14/errorcode"
	"github.com/EastWest14/gAssert"
	"net/http"
	"poemXML/poemserver/handlers/extractor"
	"poemXML/poemserver/poemlist"
	"poemXML/poemserver/poemstore"
)

const (
	DATA_ACCESS_FAILED_USER_MESSAGE = "Problem accessing poem data"
	UNKNOWN_POEM_USER_MESSAGE       = "Poem not found"
	UNKNOWN_ERROR_USER_MESSAGE      = "Unknown server error"
)

type Handlers struct {
	poemStore           PoemStoreT
	requestVarExtractor requestVarExtractorT
}

func NewHandlersInstance() *Handlers {
	return &Handlers{requestVarExtractor: extractor.New()}
}

var _ requestVarExtractorT = extractor.New()

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

func (h *Handlers) PoemReader(w http.ResponseWriter, r *http.Request) {
	poemId, ok := h.requestVarExtractor.ValueForKey(r, "poem_id")
	gAssert.Assert(ok, "Poem id not passed in - error in API processing")

	if poemId == "" {
		//Makes program compile
	}

	http.Error(w, UNKNOWN_POEM_USER_MESSAGE, 404)
}

type PoemStoreT interface {
	GetAllPoems() (list *poemlist.PoemList, err *errorcode.Errorcode)
}

//Used for unit test mocking
type requestVarExtractorT interface {
	ValueForKey(r *http.Request, key string) (value string, ok bool)
}
