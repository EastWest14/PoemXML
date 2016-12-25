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
	return server
}

func (s *Server) SetHandlers(h *handlers.Handlers) {
	s.handlers = h
}

func (s *Server) Start() {
	s.configureRouter()
	http.Handle("/", s.router)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) configureRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/poems", s.handlers.DefaultHandler)
	router.HandleFunc("/poems/list", s.handlers.PoemListHandler)
	router.HandleFunc("/poems/{poem_id}", s.handlers.PoemReader)

	s.router = router
}
