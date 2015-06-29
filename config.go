package main

import (
	"flag"
)

type Config struct {
	port string
	allowedHosts string
}

var config *Config = &Config{}

func ParseFlags() (bool) {
	flag.StringVar(&config.port, "port", ":8000", "Port to listen on")
	flag.StringVar(&config.allowedHosts, "allowed_hosts", "localhost", "Comma delimited list of acceptable host names")

	flag.Parse()
	return true
}

