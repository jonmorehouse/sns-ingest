package main

import (
	"log"
	"net/http"
)

type Server struct {
	middleware *Middleware
}

func NewServer() (*Server) {
	server := Server{}
	server.middleware = NewMiddleware()

	return &server
}

func temp(w http.ResponseWriter, r *http.Request) {
	// will be replaced by a message handler in the future
	w.Write([]byte("ok"))
}

func (s *Server) serve() {
	finalHandler := http.HandlerFunc(temp)

	err := http.ListenAndServe(config.port, s.middleware.handler(finalHandler))
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}

