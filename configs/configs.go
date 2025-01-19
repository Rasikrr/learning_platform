package configs

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

const (
	fileName = "./configs/config.toml"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Postgres PostgresConfig `toml:"postgres"`
	Redis    RedisConfig    `toml:"redis"`
	Mail     MailConfig
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
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
	DSN                 string        `toml:"dsn"`
}

type MailConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	From     string `toml:"from"`
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
	config.Postgres.DSN = os.Getenv("POSTGRES_DSN")

	config.Mail.Host = os.Getenv("MAIL_HOST")
	mailPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return Config{}, err
	}
	config.Mail.Port = mailPort
	config.Mail.User = os.Getenv("MAIL_USER")
	config.Mail.Password = os.Getenv("MAIL_PASSWORD")
	config.Mail.From = os.Getenv("MAIL_FROM")
	return config, nil
}
