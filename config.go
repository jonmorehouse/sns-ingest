package main

import (
	"log"
	"flag"
	"strings"
	"regexp"
)

type Config struct {
	port string
	allowedHostsString string
	allowedHosts []*regexp.Regexp
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

func ParseFlags() { 
	flag.StringVar(&config.port, "port", ":8000", "Port to listen on")
	flag.StringVar(&config.allowedHostsString, "allowed_hosts", "localhost", "Comma delimited list of acceptable host names")

	flag.Parse()
	config.parseAllowedHosts()
}

