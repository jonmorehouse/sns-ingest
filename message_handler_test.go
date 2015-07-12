package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestBuildsCorrectMessageType(t *testing.T) {
	request, _ := http.NewRequest("POST", "/", nil)
	messageHandler := &MessageHandler{}

	request.Header.Set("x-amz-sns-message-type", "Notification")
	message, err := messageHandler.build(request)

	assert.NotNil(t, message)
	assert.Nil(t, err)
	assert.IsType(t, message, &Notification{})

	request.Header.Set("x-amz-sns-message-type", "SubscriptionConfirmation")
	message, err = messageHandler.build(request)

	assert.NotNil(t, message)
	assert.Nil(t, err)
	assert.IsType(t, message, &Subscription{})

	request.Header.Set("x-amz-sns-message-type", "UnsubscribeConfirmation")
	message, err = messageHandler.build(request)

	assert.NotNil(t, message)
	assert.Nil(t, err)
	assert.IsType(t, message, &Unsubscription{})

	request.Header.Set("x-amz-sns-message-type", "null")
	message, err = messageHandler.build(request)

	assert.NotNil(t, err)
	assert.Nil(t, message)
}

type MockMessageHandler struct {
	mock.Mock
}

func (m MockMessageHandler) successHandler(rw http.ResponseWriter) {
	m.Called(rw)
}

func (m *MockMessageHandler) errorHandler(rw http.ResponseWriter, err error) {
	m.Called(rw, err)
}

func (m *MockMessageHandler) build(r *http.Request) (Message, error) {
	args := m.Called(r)
	return args.Get(0).(Message), args.Error(1)
}

type MockMessage struct {
	mock.Mock
}

func (m *MockMessage) handle() (error) {
	args := m.Called()
	return args.Error(0)
}

func (m *MockMessage) verify() (error) {
	args := m.Called()
	return args.Error(0)
}

func TestServerHandlesBuildError(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)
	mockError := fmt.Errorf("mock error")
	mockMessage := &MockMessage{}

	messageHandler := MockMessageHandler{}
	messageHandler.On("build", request).Return(mockMessage, mockError)
	messageHandler.On("errorHandler", responseWriter, mockError).Return()

	messageServer := HttpMessageServer{handler: &messageHandler}
	messageServer.ServeHTTP(responseWriter, request)

	messageHandler.AssertExpectations(t)
	mockMessage.AssertExpectations(t)
}

func TestServerCallsMessageHandlers(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)
	
	mockMessage := &MockMessage{}
	mockMessage.On("handle").Return(nil)
	mockMessage.On("verify").Return(nil)

	mockHandler := MockMessageHandler{}
	mockHandler.On("build", request).Return(mockMessage, nil)
	mockHandler.On("successHandler", request).Return()

	messageServer := HttpMessageServer{handler: &mockHandler}
	messageServer.ServeHTTP(responseWriter, request)

	mockHandler.AssertExpectations(t)
	mockMessage.AssertExpectations(t)
}

func TestServerHandlersMessageHandlerError(t *testing.T) {

}

