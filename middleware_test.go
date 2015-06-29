package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewMiddleware(t *testing.T) {
	middleware := NewMiddleware()

	assert.NotNil(t, middleware.validators)
}
