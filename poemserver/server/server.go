package server

import (
	"fmt"
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
	fmt.Println("Server started.")
	err := http.ListenAndServe(":8183", nil)
	if err != nil {
		fmt.Println("Server error: " + err.Error())
		panic(err.Error())
	}
}

func (s *Server) configureRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", s.handlers.DefaultHandler)
	router.HandleFunc("/poems", s.handlers.PoemListHandler)
	router.HandleFunc("/poems/{poem_id}", s.handlers.PoemReader)

	s.router = router
}
