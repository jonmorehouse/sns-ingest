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

func (m *MockMessageHandler) successHandler(rw http.ResponseWriter) {
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

func TestServerHandlesMessagesAndErrors(t *testing.T) {
	mockError := fmt.Errorf("mock error")
	var testCases = []struct {
		verifyReturn error
		handleReturn error
		handleCalled bool
		shouldSucceed bool
	}{
		{nil, mockError, true, false},
		{mockError, nil, false, false},
		{nil, nil, true, true},
	}

	for _, testCase := range testCases {
		responseWriter := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/", nil)
		mockMessage := &MockMessage{}
		mockHandler := MockMessageHandler{}
		messageServer := HttpMessageServer{handler: &mockHandler}

		mockMessage.On("verify").Return(testCase.verifyReturn)
		// handle isn't called if verify errs
		if testCase.handleCalled {
			mockMessage.On("handle").Return(testCase.handleReturn)
		}

		mockHandler.On("build", request).Return(mockMessage, nil)
		if testCase.shouldSucceed {
			mockHandler.On("successHandler", responseWriter).Return()
		} else {
			mockHandler.On("errorHandler", responseWriter, mockError).Return()
		}
		
		messageServer.ServeHTTP(responseWriter, request)
		mockMessage.AssertExpectations(t) 
		mockHandler.AssertExpectations(t)
	}
}

func TestHandlerSuccessHandler(t *testing.T) {
	messageHandler := MessageHandler{}
	responseWriter := httptest.NewRecorder()

	messageHandler.successHandler(responseWriter)

	assert.Equal(t, responseWriter.Code, 200)
	assert.Equal(t, responseWriter.Body.String(), "")
}

func TestHandlerErrorHandler(t *testing.T) {
	mockError := fmt.Errorf("error")
	messageHandler := MessageHandler{}
	responseWriter := httptest.NewRecorder()

	messageHandler.errorHandler(responseWriter, mockError)
	assert.Equal(t, responseWriter.Code, 400)
	assert.Equal(t, responseWriter.Body.String(), mockError.Error())
}

