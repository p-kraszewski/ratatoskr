package config

import (
	"github.com/creasty/defaults"
)

type Config struct {
	WebUI     string `default:"127.0.0.1:8000"`
	WireGuard string `default:"wg0"`
}

func defConfig() *Config {
	defCfg := &Config{}
	if err := defaults.Set(defCfg); err != nil {
		panic(err)
	}
	return defCfg
}
