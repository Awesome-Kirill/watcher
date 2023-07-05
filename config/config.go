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
	DebugMod      bool   `envconfig:"debug" default:"true"`

	Timeout time.Duration `envconfig:"timeout" default:"5s"`
	TTL     time.Duration `envconfig:"ttl" default:"20s"`
}

func New() *Config {
	var config Config
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
