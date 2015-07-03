package main

import (
	"regexp"
	"net/http"
	"testing"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetupMiddlewareTest() (*httptest.ResponseRecorder, *http.Request, *Middleware) {
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

	r.URL.Host = "localhost"
	r.Host = "localhost"
	err := m.hostValidator(rw, r)
	assert.Nil(t, err)

	r.Host = "not a valid host"
	err = m.hostValidator(rw, r)
	assert.NotNil(t, err)

	r.Host = "localhost"
	r.URL.Host = "not a valid host" 
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

type MockValidator struct {
	mock.Mock 
}

func (m *MockValidator) validator(rw http.ResponseWriter, r *http.Request) (error) {
	args := m.Called(rw, r)
	return args.Error(0)
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.Called(rw, r)
}

func TestHandlerSuccess(t *testing.T) {
	rw, r, m := SetupMiddlewareTest()
	validator := new(MockValidator)
	validator.On("validator", rw, r).Return(nil)

	next := new(MockHandler)
	next.On("ServeHTTP", rw, r).Return()

	m.validators = []ValidatorFunc{validator.validator}
	handler := m.handler(next)
	handler.ServeHTTP(rw, r)
	
	validator.AssertExpectations(t)
	next.AssertExpectations(t)
}

func TestHandlerFailure(t *testing.T) {
	rw, r, m := SetupMiddlewareTest()
	validator := new(MockValidator)
	validator.On("validator", rw, r).Return(fmt.Errorf("mock error"))

	next := new(MockHandler)
	m.validators = []ValidatorFunc{validator.validator}
	handler := m.handler(next)
	handler.ServeHTTP(rw, r)

	validator.AssertExpectations(t)
	next.AssertExpectations(t)
	assert.Equal(t, rw.Code, 400)
	assert.Equal(t, rw.Body.String(), "mock error")
}

