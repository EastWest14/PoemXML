package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"poemXML/poemserver/handlers"
)

type Server struct {
	router   *mux.Router
	handlers *handlers.Handlers
}

func NewServer() *Server {
	server := &Server{}
	server.configureRouter()
	return server
}

func (s *Server) SetHandlers(h *handlers.Handlers) {
	s.handlers = h
}

func (s *Server) Start() {
	http.Handle("/", s.router)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) configureRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/poems", s.handlers.DefaultHandler).Methods("GET")
	router.HandleFunc("/poems/list", s.handlers.PoemListHandler)

	s.router = router
}
