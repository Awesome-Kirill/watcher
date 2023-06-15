package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Timeout       time.Duration `envconfig:"timeout" default:"5s"`
	TTL           time.Duration `envconfig:"ttl" default:"20s"`
	ServerAddress string        `envconfig:"address" default:":1323"`
	PatchFile     string        `envconfig:"file" default:"site.txt"`
}

func New() *Config {
	var config Config
	err := envconfig.Process("todo", &config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
