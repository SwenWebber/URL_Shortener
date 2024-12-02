package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Port     string
	MongoURI string
	DBName   string
	BaseUrl  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("couldnt load .env file:%v", err)
	}

	return &Config{
		Port:     os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGODB_URI"),
		DBName:   os.Getenv("DB_NAME="),
		BaseUrl:  os.Getenv("http://localhost:8080"),
	}, nil
}
