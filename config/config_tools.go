package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

func Load(f string) (*Config, error) {
	config := defConfig()

	_, err := toml.DecodeFile(f, config)
	return config, err
}

func (cfg *Config) Save(f string) error {
	fh, err := os.Create(f)
	if err != nil {
		return err
	}
	defer fh.Close()

	return toml.NewEncoder(fh).Encode(cfg)
}
