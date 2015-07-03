package main

import (
	"log"
	"flag"
	"strings"
	"regexp"
)

type BasicAuthUser struct {
	username string
	password string
}

func (u *BasicAuthUser) verify(username string, password string) (bool) {

	if u.username == username && u.password == password {
		return true
	}

	return false
}

type Config struct {
	port string
	allowedHostsString string
	allowedHosts []*regexp.Regexp
	users []BasicAuthUser
	usersString string
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

func (c *Config) parseUsers() {
	users := strings.Split(config.usersString, ",")

	for _, user := range users {
		pieces := strings.Split(user, ":")
		basicAuthUser := BasicAuthUser{username: pieces[0], password: pieces[1]}

		config.users = append(config.users, basicAuthUser)
	}
}

func ParseFlags() {
	flag.StringVar(&config.port, "port", ":8000", "Port to listen on")
	flag.StringVar(&config.allowedHostsString, "allowed_hosts", "localhost", "Comma delimited list of acceptable host names")
	flag.StringVar(&config.usersString, "users", "", "comma delimitied string of \"username:password\" pairs")

	flag.Parse()
	config.parseAllowedHosts()
	config.parseUsers()
}



