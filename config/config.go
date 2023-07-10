package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerAddress string `envconfig:"WATCHER_PORT" default:":8080"`
	PatchFile     string `envconfig:"WATCHER_FILE" default:"site.txt"`
	AdminKey      string `envconfig:"WATCHER_ADMIN_KEY" default:"any-secret"`
	DebugMod      bool   `envconfig:"WATCHER_DEBUG" default:"true"`

	Timeout time.Duration `envconfig:"WATCHER_TIMEOUT" default:"5s"`
	TTL     time.Duration `envconfig:"WATCHER_TTL" default:"20s"`
}

func New() *Config {
	var config Config
	err := envconfig.Process("app", &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
