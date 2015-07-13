package main

import (
	"net/http"
	"errors"
	"encoding/json"
)

type HttpMessageHandler interface {
	successHandler(http.ResponseWriter)
	errorHandler(http.ResponseWriter, error)
	build(*http.Request) (Message, error)
}

type HttpMessageServer struct {
	handler HttpMessageHandler
}

func NewHttpMessageServer() (*HttpMessageServer) {
	return &HttpMessageServer{
		handler: &MessageHandler{},
	}
}

func (m *HttpMessageServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	message, err := m.handler.build(r)
	if err != nil {
		m.handler.errorHandler(rw, err)
		return
	}

	handlers := []func()(error){message.verify, message.handle}
	for _, handler := range handlers {
		err = handler()
		if err != nil {
			m.handler.errorHandler(rw, err)
			return
		}
	}

	m.handler.successHandler(rw)
}

type MessageHandler struct {}

func (m MessageHandler) successHandler(rw http.ResponseWriter) {
	rw.WriteHeader(200)
}

func (m MessageHandler) errorHandler(rw http.ResponseWriter, err error) {
	rw.WriteHeader(400)
	rw.Write([]byte(err.Error()))
}

func (mh MessageHandler) build(r *http.Request) (Message, error) {
	var err error
	var message Message

	switch messageType := r.Header.Get("x-amz-sns-message-type"); messageType {
		default:
			err = errors.New("Invalid x-amz-sns-message-type header")
		case "Notification":
			message = &Notification{}
		case "SubscriptionConfirmation":
			message = &Subscription{}
		case "UnsubscribeConfirmation":
			message = &Unsubscription{}
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(message)

	return message, err
}

