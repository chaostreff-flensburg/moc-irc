package config

import (
	env "github.com/Netflix/go-env"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Endpoint string `env:"API_ENDPOINT"`
	Addr     string `env:"IRC_ADDRESS"`
	Nick     string `env:"IRC_NICK"`
	Pass     string `env:"IRC_PASSWORD"`
	User     string `env:"IRC_USER"`
	Name     string `env:"IRC_FULLNAME"`
	Channel  string `env:"IRC_CHANNEL"`
}

// ReadConfig from env
func ReadConfig() *Config {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	if config.Endpoint == "" {
		log.Fatal("Need API_ENDPOINT env var")
	}

	if config.Addr == "" {
		log.Fatal("Need IRC_ADDRESS env var")
	}

	if config.Nick == "" {
		log.Fatal("Need IRC_NICK env var")
	}

	if config.User == "" {
		log.Fatal("Need IRC_USER env var")
	}

	if config.Name == "" {
		log.Fatal("Need IRC_FULLNAME env var")
	}

	if config.Channel == "" {
		log.Fatal("Need IRC_CHANNEL env var")
	}

	return &config
}
