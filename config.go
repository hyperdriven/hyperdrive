package hyperdrive

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

var conf Config

func init() {
	var err error
	conf, err = NewConfig()
	if err != nil {
		log.Fatalf("Config could not be initalized: %v", err)
	}
}

// Config holds configuration values from the environment, with sane defaults
// (where possible). Required configuration will throw a Fatal error if they
// are missing.
type Config struct {
	Port            int    `env:"PORT" envDefault:"5000"`
	Env             string `env:"HYPERDRIVE_ENV" envDefault:"development"`
	GzipLevel       int    `env:"GZIP_LEVEL" envDefault:"-1"`
	CorsEnabled     bool   `env:"CORS_ENABLED" envDefault:"true"`
	CorsOrigins     string `env:"CORS_ORIGINS" envDefault:"*"`
	CorsHeaders     string `env:"CORS_HEADERS" envDefault:""`
	CorsCredentials bool   `env:"CORS_CREDENTIALS" envDefault:"true"`
}

// GetPort returns the formatted value of config.Port, for use by the
// hyperdrive server, e.g. ":5000".
func (c *Config) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

// NewConfig returns an instance of config, with values loaded from ENV vars.
func NewConfig() (Config, error) {
	c := Config{}
	err := env.Parse(&c)
	return c, err
}
