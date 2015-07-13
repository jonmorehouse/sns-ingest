package main

import (
	"fmt"
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"os"
)

func GetSubscriptionFixtures() []Subscription {
	subscriptions := make([]Subscription, 0)
	file, err := os.Open("./fixtures/subscriptions.json")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&subscriptions)

	if err != nil {
		log.Fatal(err)
	}
	return subscriptions
}

func TestSubscriptionVerifyKeys(t *testing.T) {
	subscriptions := GetSubscriptionFixtures()
	testCases := []struct {
		subscription Subscription
		shouldPass bool
	}{
		{subscriptions[0], true},
		{subscriptions[1], false},
		{subscriptions[2], false},
	}

	for _, testCase := range testCases {
		fmt.Println(testCase.subscription)
		err := testCase.subscription.verifyKeys()
		if testCase.shouldPass {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "Invalid JSON")
		}
	}
}


