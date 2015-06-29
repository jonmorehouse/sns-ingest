package main

import (
	"net/http"
)

type ValidatorFunc func(http.ResponseWriter, *http.Request) (*error)

type Middleware struct {
	validators []ValidatorFunc
}

func NewMiddleware() (*Middleware) {
	middleware := Middleware{}
	middleware.validators = []ValidatorFunc{}

	return &middleware
}

func (m *Middleware) handler(next http.Handler) (http.Handler) {

	return next
}


