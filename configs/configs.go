package configs

import "github.com/BurntSushi/toml"

const (
	fileName = "./configs/cfg.toml"
)

type Config struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func Parse() (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(fileName, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
