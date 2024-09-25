package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	USER_PORT           string
	NATIONALITY_SERVICE string

	DB_PORT     string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() Config {
	//if err := godotenv.Load(".env"); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	config := Config{}

	config.USER_PORT = cast.ToString(coalesce("USER_PORT", ":50050"))
	config.NATIONALITY_SERVICE = cast.ToString(coalesce("NATIONALITY_SERVICE", ":7080"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "dodi"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "cm"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
