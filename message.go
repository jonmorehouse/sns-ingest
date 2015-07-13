package main

import (
	"errors"
)

type Message interface {
	// verify encapsulates the signature creation 
	// returns the output of the signature verification step
	verify() (error)

	// handle takes the proper action
	// either feeds to nsq/webcallback or handles subscription
	handle() (error)
}


type Subscription struct {
	Type string `json:"Type"`
	MessageID string `json:"MessageID"`
	Token string `json:"Token"`
	TopicArn string `json:"TopicArn"`

	SubscribeURL string `json:"SubscribeURL"`
	Timestamp string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature string `json:"Signature"`
	SigningCertURL string `json:"SigningCertURL"`
}

func (*Subscription) verify() (error) {

	return nil
}

func (*Subscription) handle() (error) {

	return nil
}

func VerifyKeys(keys *[]*string) (error) {
	
	for _, key := range *keys {
		if *key == "" {
			err := errors.New("Invalid JSON")
			return err
		}
	}

	return nil
}

func (s *Subscription) verifyKeys() (error) {
	keys := []*string{
		&s.Type,		
		&s.MessageID,
		&s.Token,
		&s.TopicArn,
		&s.SubscribeURL,
		&s.Timestamp,
		&s.SignatureVersion,
		&s.Signature,
		&s.SigningCertURL,
	}

	return VerifyKeys(&keys)
}

type Unsubscription struct {
	Type string `json:"Type"`
	MessageID string `json:"MessageID"`
	Message string `json:"Message"`
	Token string `json:"Token"`
	TopicArn string `json:"TopicArn"`
	Subject string `json:"Subject"`
	Timestamp string `json:"Timestamp"`
	UnSubscribeURL string `json:"UnsubscribeURL"`
}

func (*Unsubscription) verify() (error) {
	return nil
}

func (*Unsubscription) handle() (error) {
	return nil
}

type Notification struct {
	Type string `json:"Type"`
	MessageID string `json:"MessageID"`
	Message string `json:"Message"`
	Token string `json:"Token"`
	TopicArn string `json:"TopicArn"`
	Subject string `json:"Subject"`
	Timestamp string `json:"Timestamp"`
	UnSubscribeURL string `json:"UnsubscribeURL"`
}

func (n *Notification) verify() (error) {
	return nil
}

func (n *Notification) handle() (error) {

	return nil
}


