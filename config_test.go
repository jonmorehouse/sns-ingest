package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseAllowedHosts(t *testing.T) {
	originalAllowedHosts := config.allowedHostsString
	config.allowedHostsString = "^localhost,^localhost$"
	config.parseAllowedHosts()

	assert.Equal(t, 2, len(config.allowedHosts))

	// basic tests to ensure that the regexp's are parsed correctly
	assert.True(t, config.allowedHosts[0].MatchString("localhost"))
	assert.True(t, config.allowedHosts[0].MatchString("localhosta"))
	assert.False(t, config.allowedHosts[0].MatchString("alocalhost"))
	assert.False(t, config.allowedHosts[1].MatchString("localhosta"))
	assert.True(t, config.allowedHosts[1].MatchString("localhost"))

	config.allowedHostsString = originalAllowedHosts
	config.parseAllowedHosts()
}
