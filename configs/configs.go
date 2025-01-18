package configs

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"os"
	"time"
)

const (
	fileName = "./configs/cfg.toml"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Postgres PostgresConfig `toml:"postgres"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type PostgresConfig struct {
	DBName              string        `toml:"db_name"`
	User                string        `toml:"user"`
	Password            string        `toml:"password"`
	Host                string        `toml:"host"`
	Port                string        `toml:"port"`
	MaxConns            int           `toml:"max_conns"`
	MinConns            int           `toml:"min_conns"`
	MaxIdleConnIdleTime time.Duration `toml:"max_idle_conn_time"`
}

func Parse() (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(fileName, &config); err != nil {
		return Config{}, err
	}
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	config.Postgres.DBName = os.Getenv("POSTGRES_DB")
	config.Postgres.User = os.Getenv("POSTGRES_USER")
	config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Postgres.Host = os.Getenv("POSTGRES_HOST")
	config.Postgres.Port = os.Getenv("POSTGRES_PORT")
	return config, nil
}
