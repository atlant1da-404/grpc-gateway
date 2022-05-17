package configs

import (
	"os"
)

type apiConfigs struct {
	HttpHost string
	HttpPort string

	GrpcHost string
	GrpcPort string
}

func API() *apiConfigs {

	return &apiConfigs{

		HttpHost: os.Getenv("httpHost"),
		HttpPort: os.Getenv("httpPort"),

		GrpcHost: os.Getenv("grpcHost"),
		GrpcPort: os.Getenv("grpcPort"),
	}
}
