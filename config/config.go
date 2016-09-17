package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/caarlos0/env.v2"
)

// Cfg default system configurations
var Cfg Config

type (
	// Config represents all system configurations
	Config struct {
		Env        string `env:"GO_ENV" envDefault:"development"`
		Port       int    `env:"PORT" envDefault:"8081"`
		APPVersion string `env:"APP_VERSION" envDefault:"1.0"`

		TokenSecret string `env:"TOKEN_SECRET"`

		UserEndpoint string `env:"ENDPOINT_USER" envDefault:"http://localhost:8081/api/user"`
	}

	// Provider ...
	Provider interface {
		Get() *Config
		IsDevelopment() bool
	}

	// EnvProvider ...
	EnvProvider struct{}
)

// Init initializes Config from env variables
func Init() *Config {
	godotenv.Load()
	env.Parse(&Cfg)

	return &Cfg
}

// Get config from memory
// Must call Init on application startup
func (cfg *EnvProvider) Get() *Config {
	return &Cfg
}

// IsDevelopment ...
func (cfg *EnvProvider) IsDevelopment() bool {
	return Cfg.Env != "production"
}
