package handlers

import (
	"poemXML/poemserver/poemstore"
)

func (h *Handlers) GetPoemStore() *poemstore.Store {
	return h.poemStore
}
