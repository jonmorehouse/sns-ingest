package main

import (
	"net/http"
)

type MessageHandler struct {}

func NewMessageHandler() (*MessageHandler) {
	return &MessageHandler{}
}

func (mh *MessageHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	message := MessageWithRequest(r)
	err := message.Validate()

	// handles the result of the messages acceptance
	if err != nil {
		// handle error
		// return 400
		http.NotFound(rw, r)
	} else {
		rw.Write([]byte("Message accepted and acknowledged"))
	}
}

func MessageWithRequest(r *http.Request) (*Message) {

	return nil
}


