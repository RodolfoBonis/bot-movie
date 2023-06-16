package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

func EnvBotToken() string {
	return GetEnv("DISCORD_TOKEN", "")
}

func EnvChannelID() string {
	return GetEnv("CHANNEL_ID", "")
}

func EnvErmaclessID() string {
	return GetEnv("ERMACLESS_ID", "")
}

func EnvLaisID() string {
	return GetEnv("LAIS_ID", "")
}

func EnvAlnWolfID() string {
	return GetEnv("ALNWOLF_ID", "")
}

func EnvJonasID() string {
	return GetEnv("JONAS_ID", "")
}

func EnvTimeZone() string {
	return GetEnv("TIME_ZONE", "America/Sao_Paulo")
}

func LoadEnvVars() {
	env := GetEnv("BOT_ENV", "development")

	if env == "production" || env == "staging" {
		log.Println("Not using .env file in production or staging.")
		return
	}

	filename := ".env." + env

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = ".env"
	}

	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal(".env file not loaded")
	}
}
