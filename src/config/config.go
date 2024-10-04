package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port        string
	DatabaseUrl string
}

func GetVariablesEnv() *EnvConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	return &EnvConfig{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASEURL"),
	}
}
