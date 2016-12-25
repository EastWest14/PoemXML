package extractor

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RequestVarExtractor struct{}

func New() *RequestVarExtractor {
	return &RequestVarExtractor{}
}

func (rve *RequestVarExtractor) ValueForKey(r *http.Request, key string) (value string, ok bool) {
	requestVars := mux.Vars(r)
	value, ok = requestVars[key]
	return value, ok
}
