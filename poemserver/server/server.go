package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"poemXML/poemserver/poemstore"
	"strings"
)

const STANDARD_RESPONSE = "Hello, world!"

type Server struct {
	poemStore *poemstore.Store
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateAndConfigureRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/poems", s.DefaultHandler).Methods("GET")
	r.HandleFunc("/poems/list", s.PoemListHandler)

	return r
}

func (s *Server) SetPoemStore(pStore *poemstore.Store) {
	s.poemStore = pStore
}

func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	fmt.Println(urlPath)

	fmt.Fprint(w, STANDARD_RESPONSE)
}

func (s *Server) PoemListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of poems")
}

func firstPartOfPath(u *url.URL) string {
	urlPath := u.Path
	pathParts := strings.Split(urlPath, "/")

	if len(pathParts) > 1 {
		return pathParts[1]
	}

	return ""
}
