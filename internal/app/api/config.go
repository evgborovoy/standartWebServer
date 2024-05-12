package api

import "github.com/evgborovoy/StandardWebServer/storage"

// General instance for API server

type Config struct {
	// Port
	BindAddr string `toml:"bind_addr"`
	// Logger
	LoggerLevel string `toml:"logger_level"`
	// Storage config
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
