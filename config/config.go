package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerAddress string `envconfig:"address" default:":8080"`
	PatchFile     string `envconfig:"file" default:"site.txt"`
	AdminKey      string `envconfig:"adminKey" default:"any-secret"`

	Timeout         time.Duration `envconfig:"timeout" default:"5s"`
	TimeoutShutdown time.Duration `envconfig:"timeoutShut" default:"15s"`
	TTL             time.Duration `envconfig:"ttl" default:"20s"`
}

func New() *Config {
	var config Config
	err := envconfig.Process("todo", &config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
