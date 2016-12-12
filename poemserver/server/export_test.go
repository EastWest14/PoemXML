package server

import (
	"poemXML/poemserver/handlers"
)

func (s *Server) GetHandlers() *handlers.Handlers {
	return s.handlers
}
