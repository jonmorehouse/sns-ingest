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

func TestParseUsers(t *testing.T) {
	originalUsersString := config.usersString
	config.usersString = "username:password,username:password2,username2:password"
	config.parseUsers()
	
	assert.Equal(t, 3, len(config.users))
	assert.Equal(t, config.users[0].username, "username")
	assert.Equal(t, config.users[1].username, "username")
	assert.Equal(t, config.users[2].username, "username2")

	assert.Equal(t, config.users[0].password, "password")
	assert.Equal(t, config.users[1].password, "password2")
	assert.Equal(t, config.users[2].password, "password")

	config.usersString = originalUsersString
}

func TestBasicAuthUserVerify(t *testing.T) {
	user := BasicAuthUser{"username", "password"}

	assert.True(t, user.verify("username", "password"))
	assert.False(t, user.verify("username", "bad password"))
}

