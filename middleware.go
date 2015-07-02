package main

import (
	"net/http"
)

type ValidatorFunc func(http.ResponseWriter, *http.Request) (*error)

type Middleware struct {
	validators []ValidatorFunc
}

func NewMiddleware() (*Middleware) {
	m := Middleware{}
	m.validators = []ValidatorFunc{m.hostnameValidator, m.basicAuthValidator, m.contentTypeValidator}

	return &m
}

func (m *Middleware) hostnameValidator(rw http.ResponseWriter, r *http.Request) (*error) {

	return nil
}

func (m *Middleware) contentTypeValidator(rw http.ResponseWriter, r *http.Request) (*error) {

	return nil
}

func (m *Middleware) basicAuthValidator(rw http.ResponseWriter, r *http.Request) (*error) {

	return nil
}

func (m *Middleware) handler(next http.Handler) (http.Handler) {


	return next
}

