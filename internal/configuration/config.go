package configuration

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerHTTP HTTP       `toml:"server"`
	Logger     LoggerConf `toml:"logger"`
}

type HTTP struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type LoggerConf struct {
	Level   string `toml:"log_level"`
	File    string `toml:"log_file"`
	IsProd  bool   `toml:"log_trace_on"`
	TraceOn bool   `toml:"log_prod_on"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
