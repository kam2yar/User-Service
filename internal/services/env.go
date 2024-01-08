package services

import (
	"github.com/joho/godotenv"
	"log"
)

var env map[string]string
var loaded = false

func Env(key string) string {
	if !loaded {
		loadEnv()
	}

	return env[key]
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	env, err = godotenv.Read()
	if err != nil {
		log.Fatalf("Error Reading .env file: %s", err)
	}

	loaded = true
}
