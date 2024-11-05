package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBPort string `env:"DB_PORT,required"`
	DBHost string `env:"DB_HOST,required"`
	DBName string `env:"DB_NAME,required"`
	DBUser string `env:"DB_USER,required"`
	DBPass string `env:"DB_PASS,required"`
}

func NewEnvConfig() *EnvConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env: %e", err)
	}

	config := &EnvConfig{}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env %s:", err)
	}

	return config
}
