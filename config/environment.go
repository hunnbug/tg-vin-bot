package config

import (
	env "github.com/joho/godotenv"
)

func LoadEnv() error {
	return env.Load()
}
