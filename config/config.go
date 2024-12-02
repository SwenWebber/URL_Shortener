package config

import (
	"os"
)

type Config struct {
	Port     string
	MongoURI string
	DBName   string
	BaseUrl  string
}

func LoadConfig() (*Config, error) {

	return &Config{
		Port:     os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGODB_URI"),
		DBName:   os.Getenv("DB_NAME"),
		BaseUrl:  os.Getenv("BASE_URL"),
	}, nil
}
