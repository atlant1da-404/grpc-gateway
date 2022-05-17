package main

import (
	"github.com/joho/godotenv"
	"grpc-gateway/configs"
	"grpc-gateway/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const env = "local.env"

func main() {

	err := godotenv.Load(env)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	cfg := configs.API()

	srv := server.NewServer(cfg.HttpHost+":"+cfg.HttpPort, cfg.GrpcHost+":"+cfg.GrpcPort)
	go srv.RunAPIServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

}
