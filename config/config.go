package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Config struct {
	Timeout   time.Duration `envconfig:"timeout" default:"5s"`
	TTL       time.Duration `envconfig:"ttl" default:"20s"`
	PatchFile string        `envconfig:"file" default:"site.txt"`
}

func New() *Config {

	var config Config
	err := envconfig.Process("todo", &config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
