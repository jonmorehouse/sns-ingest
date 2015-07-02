package main

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

import "fmt"

func TestParseAllowedHosts(t *testing.T) {
	originalAllowedHosts := config.allowedHostsString
	config.allowedHostsString = "localhost"//,.*localhost"
	config.parseAllowedHosts()

	res := config.allowedHosts[0].Match([]byte("non-localhost"))
	if res == true {
		fmt.Println("match")
	}

	config.allowedHostsString = originalAllowedHosts
	config.parseAllowedHosts()

}
