package hyperdrive

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

// Config holds configuration values from the environment, with sane defaults
// (where possible). Required configuration will throw a Fatal error if they
// are missing.
type Config struct {
	Port      int    `env:"PORT" envDefault:"5000"`
	Env       string `env:"HYPERDRIVE_ENV" envDefault:"development"`
	GzipLevel int    `env:"GZIP_LEVEL" envDefault:"-1"`
}

// GetPort returns the formatted value of config.Port, for use by the
// hyperdrive server, e.g. ":5000".
func (c *Config) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

// NewConfig returns an instance of config, with values loaded from ENV vars.
func NewConfig() Config {
	c := Config{}
	err := env.Parse(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
