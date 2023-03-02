package config

import (
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func CorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return config
}
