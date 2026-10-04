package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() Config {
	godotenv.Load()
	Port := os.Getenv("PORT")
	if Port == "" {
		panic("PORT is required")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}
	return Config{
		Port,
		env,
	}
}
