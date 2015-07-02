package main

import (
	"net/http"
	"errors"
)

type ValidatorFunc func(http.ResponseWriter, *http.Request) (error)

type Middleware struct {
	validators []ValidatorFunc
}

func NewMiddleware() (*Middleware) {
	m := Middleware{}
	m.validators = []ValidatorFunc{
		m.httpMethodValidator, 
		m.contentTypeValidator,
		m.hostnameValidator, 
		m.basicAuthValidator, 
		m.uriValidator,
	}

	return &m
}

func (m *Middleware) httpMethodValidator(rw http.ResponseWriter, r *http.Request) (error) {
	if r.Method != "POST" {
		return errors.New("Invalid message type")
	}
	return nil
}

func (m *Middleware) hostnameValidator(rw http.ResponseWriter, r *http.Request) (error) {

	return nil
}

func (m *Middleware) contentTypeValidator(rw http.ResponseWriter, r *http.Request) (error) {

	return nil
}

func (m *Middleware) basicAuthValidator(rw http.ResponseWriter, r *http.Request) (error) {

	return nil
}

func (m *Middleware) uriValidator(rw http.ResponseWriter, r *http.Request) (error) {
	// check that uri can be resolved to an nsq queue name

	return nil
}

func (m *Middleware) handler(next http.Handler) (http.Handler) {

	return next
}

