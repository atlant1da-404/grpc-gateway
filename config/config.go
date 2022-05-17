package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	env              = "local.env"
	defaultSocketOne = "localhost:8080"
	defaultSocketTwo = "localhost:8081"
)

func LoadSocketConfig() []string {

	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	socketOne, socketTwo := os.Getenv("addressOne"), os.Getenv("addressTwo")
	if socketOne == "" && socketTwo == "" {
		socketOne = defaultSocketOne
		socketTwo = defaultSocketTwo
	}

	return []string{socketOne, socketTwo}
}
