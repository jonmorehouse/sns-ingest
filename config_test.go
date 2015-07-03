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

	user, ok := config.users[0].(BasicAuthUser)
	assert.True(t, ok)
	assert.Equal(t, user.username, "username")
	assert.Equal(t, user.password, "password")

	user, ok = config.users[1].(BasicAuthUser)
	assert.True(t, ok)
	assert.Equal(t, user.username, "username")
	assert.Equal(t, user.password, "password2")

	user, ok = config.users[2].(BasicAuthUser)
	assert.True(t, ok)
	assert.Equal(t, user.username, "username2")
	assert.Equal(t, user.password, "password")

	config.usersString = originalUsersString
}

func TestBasicAuthUserVerify(t *testing.T) {
	user := BasicAuthUser{"username", "password"}

	assert.True(t, user.verify("username", "password"))
	assert.False(t, user.verify("username", "bad password"))
}

func TestParseAllowedQueues(t *testing.T) {
	assert.Equal(t, len(config.allowedQueues), 0)

	config.allowedQueuesString = "queue_1,queue_2"
	config.parseAllowedQueues()
	assert.Equal(t, len(config.allowedQueues), 2)
	assert.Equal(t, config.allowedQueues[0], "queue_1")
	assert.Equal(t, config.allowedQueues[1], "queue_2")

	config.allowedQueuesString = ""
	config.parseAllowedQueues()
}

