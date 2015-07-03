package main

import (
	"regexp"
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetupMiddlewareTest() (http.ResponseWriter, *http.Request, *Middleware) {
	responseWriter := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)
	middleware := NewMiddleware()

	return responseWriter, request, middleware
}

func TestHttpMethodValidator(t *testing.T) {
	rw, r, m := SetupMiddlewareTest()

	invalidMethods := []string{"GET", "PUT", "PATCH", "HEAD"}
	for _, method := range invalidMethods {
		r.Method = method
		err := m.httpMethodValidator(rw, r)
		assert.NotNil(t, err)
	}

	r.Method = "POST"
	assert.Nil(t, m.httpMethodValidator(rw, r))
}

func TestHostValidator(t *testing.T) {
	rw, r, m := SetupMiddlewareTest()
	originalAllowedHosts := config.allowedHosts

	regex, _ := regexp.Compile("^localhost$")
	config.allowedHosts = []*regexp.Regexp{regex}

	r.Host = "localhost"
	err := m.hostValidator(rw, r)
	assert.Nil(t, err)

	r.Host = "not a valid host"
	err = m.hostValidator(rw, r)
	assert.NotNil(t, err)

	config.allowedHosts = originalAllowedHosts
}

func TestSnsHeaderValidator(t *testing.T) {
	rw, r, m := SetupMiddlewareTest()

	r.Header.Set("Content-Type", "content-type")
	err := m.snsHeaderValidator(rw, r)
	assert.Equal(t, err.Error(), "Invalid Content-Type")

	r.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	r.Header.Set("User-Agent", "user-agent")
	err = m.snsHeaderValidator(rw, r)
	assert.Equal(t, err.Error(), "Invalid User-Agent")

	r.Header.Set("User-Agent", "Amazon Simple Notification Service Agent")

	snsHeaders := []string {
		"x-amz-sns-message-type",
		"x-amz-sns-message-id",
		"x-amz-sns-topic-arn",
	}

	for _, header := range snsHeaders {
		err = m.snsHeaderValidator(rw, r)
		assert.Equal(t, err.Error(), "Missing header " + header)
		r.Header.Set(header, "value")
	}

	err = m.snsHeaderValidator(rw, r)
	assert.Nil(t, err)
}

type BasicAuthMock struct {
	mock.Mock
}

func (b *BasicAuthMock) verify(username, password string) (bool) {
	args := b.Called(username, password)
	return args.Bool(0)
}

func TestBasicAuthValidator(t *testing.T) {
	mock := new(BasicAuthMock)
	mock.On("verify", "username", "password").Return(true)
	mock.On("verify", "bad_username", "bad_password").Return(false)

	originalUsers := config.users
	config.users = []User{mock}

	rw, r, m := SetupMiddlewareTest()
	r.SetBasicAuth("username", "password")
	err := m.basicAuthValidator(rw, r)
	assert.Nil(t, err)

	r.SetBasicAuth("bad_username", "bad_password")
	err = m.basicAuthValidator(rw, r)

	mock.AssertExpectations(t)
	config.users = originalUsers
}

func TestUriValidator(t *testing.T) {
	rw, r, m := SetupMiddlewareTest()
	config.allowedQueues = map[string]bool{
		"queue": true,
		"queue_2": true,
		"queue_3": true,
	}

	r, _ = http.NewRequest("POST", "http://localhost/queue/queue_2/queue_3", nil)
	err := m.uriValidator(rw, r)

	assert.Nil(t, err)

	r, _ = http.NewRequest("POST", "http://localhost/queue/bad_queue/bad_queue", nil)
	err = m.uriValidator(rw, r)
	assert.NotNil(t, err)
}

