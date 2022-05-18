package main

import (
	"github.com/joho/godotenv"
	"grpc-gateway/configs"
	"grpc-gateway/internal/repository"
	"grpc-gateway/server"
	"log"
)

const env = "local.env"

func main() {

	err := godotenv.Load(env)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	cfg := configs.Configs()
	repository.NewRedis(cfg.DbHost + ":" + cfg.DbPort)

	srv := server.NewServer(cfg.HttpHost+":"+cfg.HttpPort, cfg.GrpcHost+":"+cfg.GrpcPort)
	srv.RunAPIServer()
}
