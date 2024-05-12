package storage

type Config struct {
	DataBaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
