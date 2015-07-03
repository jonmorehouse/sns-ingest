package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"regexp"
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

func TestBasicAuthValidator(t *testing.T) {
	originalUsers := config.users
	config.users = []BasicAuthUser{{"username", "password"}}

	assert.Equal(t, config.users[0].username, "username")

	config.users = originalUsers
}

func TestUriValidator(t *testing.T) {

}


