package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := NewServer()

	assert.NotNil(t, server.middleware)
}


