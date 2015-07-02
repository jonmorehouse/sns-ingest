package main

import (
	"log"
	"net/http"
)

type Server struct {
	middleware *Middleware
	handler *MessageHandler
}

func NewServer() (*Server) {
	server := Server{}
	server.middleware = NewMiddleware()
	server.handler = NewMessageHandler()

	return &server
}

func (s *Server) serve() {
	err := http.ListenAndServe(config.port, s.middleware.handler(s.handler))

	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

