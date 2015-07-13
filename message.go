package main

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
	MessageID, 
	Token,
	TopicArn,
	SubscribeURL,
	Timestamp,
	SignatureVersion,
	Signature,
	SigningCertURL string
	url string
}

func (*Subscription) verify() (error) {

	return nil
}

func (*Subscription) handle() (error) {

	return nil
}

type Unsubscription struct {
	Message,
	MessageId,
	Subject,
	Timestamp,
	TopicArn,
	UnSubscribeURL,
	Type string	    
}

func (*Unsubscription) verify() (error) {

	return nil
}

func (*Unsubscription) handle() (error) {

	return nil
}

type Notification struct {
	TypeA string `json:"Type"`
	Message,
	MessageId,
	Subject,
	Timestamp,
	TopicArn,
	UnSubscribeURL,
	Type string	    
}

func (n *Notification) verify() (error) {
	return nil
}

func (n *Notification) handle() (error) {

	return nil
}


