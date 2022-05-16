package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	env = "local.env"
)

func LoadSocketConfig() (string, string) {

	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf(err.Error())
		return "", ""
	}

	return os.Getenv("addressOne"), os.Getenv("addressTwo")
}
