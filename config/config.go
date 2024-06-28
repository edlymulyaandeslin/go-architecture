package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	Driver   string
}

type SecurityConfig struct {
}
type AppConfig struct {
	AppPort string
}

type Config struct {
	DbConfig
	AppConfig
	SecurityConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.DbConfig = DbConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	c.AppConfig = AppConfig{
		AppPort: os.Getenv("APP_PORT"),
	}

	if c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.DbName == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.AppConfig.AppPort == "" {
		return errors.New("enviroment is empty")
	}
	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.readConfig(); err != nil {
		return nil, err
	}
	return config, nil
}
