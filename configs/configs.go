package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	devFileConfig  = "./configs/config.toml"
	prodFileConfig = "./configs/config_prod.toml"
)

const (
	Development = "dev"
	Production  = "prod"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Postgres PostgresConfig `toml:"postgres"`
	Redis    RedisConfig    `toml:"redis"`
	Auth     AuthConfig     `toml:"auth"`
	Mail     MailConfig
	JDoodle  JDoodleConfig
}

type AuthConfig struct {
	Secret               string        `toml:"secret"`
	AccessTokenLifeTime  time.Duration `toml:"access_token_lifetime"`
	RefreshTokenLifeTime time.Duration `toml:"refresh_token_lifetime"`
}
type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
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

type JDoodleConfig struct {
	URL          string
	ClientID     string
	ClientSecret string
}

func Parse() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file")
	}

	env := os.Getenv("ENV")

	var config Config
	switch env {
	case Development:
		if _, err := toml.DecodeFile(devFileConfig, &config); err != nil {
			return Config{}, err
		}
	case Production:
		if _, err := toml.DecodeFile(prodFileConfig, &config); err != nil {
			return Config{}, err
		}
	default:
		return Config{}, fmt.Errorf("invalid environment: %s", env)
	}
	log.Printf("Initializing config for env: %s\n", env)

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

	config.JDoodle.URL = os.Getenv("JDOODLE_URL")
	config.JDoodle.ClientID = os.Getenv("JDOODLE_CLIENT_ID")
	config.JDoodle.ClientSecret = os.Getenv("JDOODLE_CLIENT_SECRET")

	config.Mail.Port = mailPort
	config.Mail.User = os.Getenv("MAIL_USER")
	config.Mail.Password = os.Getenv("MAIL_PASSWORD")
	config.Mail.From = os.Getenv("MAIL_FROM")

	config.Redis.User = os.Getenv("REDIS_USER")
	config.Redis.Password = os.Getenv("REDIS_PASSWORD")

	config.Auth.Secret = os.Getenv("JWT_SECRET")

	return config, nil
}
