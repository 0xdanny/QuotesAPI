package server

import (
	"github.com/0xdanny/QuotesAPI/server/database"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     database.QuotesDB
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.initRoutes()
	return s
}

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/", s.HomeHandler()).Methods("GET")
	s.Router.HandleFunc("/api/quotes", s.GetQuotesHandler()).Methods("GET")
	s.Router.HandleFunc("/api/quotes", s.CreateQuoteHandler()).Methods("POST")
	s.Router.HandleFunc("/api/quotes/{id}", s.DeleteQuoteHandler()).Methods("DELETE")
	s.Router.HandleFunc("/api/quotes", s.UpdateQuoteHandler()).Methods("PUT")
}
