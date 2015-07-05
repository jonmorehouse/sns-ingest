package main

import (
	"net/http"
)

type MessageHandler struct {}

func NewMessageHandler() (*MessageHandler) {
	return &MessageHandler{}
}

func (mh *MessageHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//message := MessageWithRequest(r)
	//err := message.Validate()
	//err := nil

	//// handles the result of the messages acceptance
	//if err != nil {
		//rw.WriteHeader(400)
		//rw.Write([]byte(err.Error()))
	//} else {
		//rw.WriteHeader(204)
	//}
}

//func MessageWithRequest(r *http.Request) (*Message) {


//}

