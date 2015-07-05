package main

const (
	SUBSCRIPTION = iota
	UNSUBSCRIPTION = iota
	NOTIFICATION = iota
)

// signature mixin is responsible for taking the signature string and verifying it 
type SignatureMixin struct {
	// keeps a certificate object alive
	certificate *Certificate,

}

func (s *SignatureMixin) downloadCertURL(certURL string) (error) {
	// this is run in a go-process behind the scenes
	// handles caching seamlessly
	//s.certificate = CertificateWithURL(certURL)
	return nil
}

func (s *SignatureMixin) signMessage(message string) (error) {
	// base64 encode message
	base64EncodedMessage := base64.base64(message)
	pubKey := s.certificate.getPublicKey()

	// this is going to be cpu bound, probably not worth splitting into a go routine
	return pubKey.sign(base64EncodedMessage)
}

func (s *SignatureMixin) verifySignature(expectedSignature, signature string) (error) {

	return nil
}


type SubscriptionNotification struct {
	Type,
	MessageID,
	Token,
	TopicArn,
	SubscribeURL,
	Timestamp,
	SignatureVersion,
	Signature,
	SigningCertURL string
	url string

	SignatureMixin
}

type MessageNotification struct {
	Message,
	MessageId,
	Subject,
	Timestamp,
	TopicArn,
	UnSubscribeURL,
	Type string	    

	SignatureMixin
}

