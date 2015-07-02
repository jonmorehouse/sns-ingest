package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
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

func TestHostNameValidator(t *testing.T) {
	//rw, r, m := SetupMiddlewareTest()

	// set host names
	// set validators

}

