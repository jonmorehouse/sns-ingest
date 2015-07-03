package main

import (
	"log"
	"flag"
	"strings"
	"regexp"
)

type User interface {
	verify(string, string) bool
}

type BasicAuthUser struct {
	username string
	password string
}

func (u BasicAuthUser) verify(username, password string) (bool) {
	if u.username == username && u.password == password {
		return true
	}

	return false
}

type Config struct {
	port string
	allowedHostsString string
	allowedHosts []*regexp.Regexp
	basicAuthEnabled bool
	users []User
	usersString string
	allowedQueuesString string
	allowedQueues map[string]bool
}

var config *Config = &Config{}

func (c *Config) parseAllowedHosts() {
	pieces := strings.Split(config.allowedHostsString, ",")

	for _, piece := range pieces {
		regex, err := regexp.Compile(piece)

		if err != nil {
			log.Fatal(err)
		} else {
			config.allowedHosts = append(config.allowedHosts, regex) 
		}
	}
}

func (c *Config) parseAllowedQueues() {
	allowedQueues := strings.Split(config.allowedQueuesString, ",")
	
	config.allowedQueues = make(map[string]bool)
	for _, queue := range allowedQueues {
		config.allowedQueues[queue] = true
	}
}

func (c *Config) parseUsers() {
	users := strings.Split(config.usersString, ",")

	for _, user := range users {
		pieces := strings.Split(user, ":")
		basicAuthUser := BasicAuthUser{username: pieces[0], password: pieces[1]}

		config.users = append(config.users, basicAuthUser)
	}

	config.basicAuthEnabled = len(config.users) > 0
}

func ParseFlags() {
	flag.StringVar(&config.port, "port", ":8000", "Port to listen on")
	flag.StringVar(&config.allowedHostsString, "allowed_hosts", "localhost", "Comma delimited list of acceptable host names")
	flag.StringVar(&config.usersString, "users", "", "comma delimitied string of \"username:password\" pairs")
	flag.StringVar(&config.allowedQueuesString, "queues", "", "comma delimited string of nsq queue names")

	flag.Parse()

	config.parseAllowedHosts()
	config.parseUsers()
	config.parseAllowedQueues()
}

