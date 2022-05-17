package configs

import (
	"os"
)

type apiConfigs struct {
	HttpHost string
	HttpPort string

	GrpcHost string
	GrpcPort string

	DbHost string
	DbPort string
}

func Configs() *apiConfigs {

	return &apiConfigs{

		HttpHost: os.Getenv("httpHost"),
		HttpPort: os.Getenv("httpPort"),

		GrpcHost: os.Getenv("grpcHost"),
		GrpcPort: os.Getenv("grpcPort"),

		DbHost: os.Getenv("dbHost"),
		DbPort: os.Getenv("dbPort"),
	}
}
