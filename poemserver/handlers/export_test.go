package handlers

import ()

func (h *Handlers) GetPoemStore() PoemStoreT {
	return h.poemStore
}

func (h *Handlers) SetRequestVarExtractor(varExtractor requestVarExtractorT) {
	h.requestVarExtractor = varExtractor
}
