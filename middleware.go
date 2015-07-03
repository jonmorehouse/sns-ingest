package main

import (
	"fmt"
	"errors"
	"net/http"
)

type ValidatorFunc func(http.ResponseWriter, *http.Request) (error)

type Middleware struct {
	validators []ValidatorFunc
}

func NewMiddleware() (*Middleware) {
	m := Middleware{}
	m.validators = []ValidatorFunc{
		m.httpMethodValidator, 
		m.snsHeaderValidator,
		m.hostValidator, 
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

func (m *Middleware) hostValidator(rw http.ResponseWriter, r *http.Request) (error) {
	var found bool
	for _, regex := range config.allowedHosts {
		found = regex.MatchString(r.Host)
		if found {
			return nil
		}
	}

	return errors.New("Invalid host name")
}

func (m *Middleware) snsHeaderValidator(rw http.ResponseWriter, r *http.Request) (error) {
	if val := r.Header.Get("Content-Type"); val != "text/plain; charset=UTF-8" {
		return errors.New("Invalid Content-Type")
	}

	if val := r.Header.Get("User-Agent"); val != "Amazon Simple Notification Service Agent" {
		return errors.New("Invalid User-Agent")
	}

	awsSNSHeaders := []string{
		"x-amz-sns-message-type",
		"x-amz-sns-message-id",
		"x-amz-sns-topic-arn",
	}

	for _, header := range awsSNSHeaders {
		val := r.Header.Get(header)
		if val == "" {
			return fmt.Errorf("Missing header %s", header)
		}
	}

	return nil
}

func (m *Middleware) basicAuthValidator(rw http.ResponseWriter, r *http.Request) (error) {


	return nil
}

func (m *Middleware) uriValidator(rw http.ResponseWriter, r *http.Request) (error) {
	// checks valid queue names

	return nil
}

func (m *Middleware) handler(next http.Handler) (http.Handler) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//var err error

		//for _, validator := range m.validators

		m.hostValidator(rw, r)
	})
}

