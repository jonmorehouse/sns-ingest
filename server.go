package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router *httprouter.Router
}

func InitServer() (*Server) {
	server := Server{}
	server.router = httprouter.New()

	return &server
}

func (s *Server) serve() {
	err := http.ListenAndServe(config.port, s.router)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

