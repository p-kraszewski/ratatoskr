package config

import (
	"github.com/BurntSushi/toml"
)

func Load(f string) (*Config, error) {
	config := defConfig()

	_, err := toml.DecodeFile(f, config)
	return config, err
}

func defConfig() *Config {
	return &Config{}
}
