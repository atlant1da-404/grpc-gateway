package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc-gateway/config"
	"grpc-gateway/internal/app"
	"grpc-gateway/pkg/gen/notes"
	"log"
	"net"
	"net/http"
)

const (
	method = "tcp"
)

func main() {

	hosts := config.LoadSocketConfig()

	go func() {

		mux := runtime.NewServeMux()
		notes.RegisterNotesHandlerServer(context.Background(), mux, app.NewNotes())
		log.Fatalln(http.ListenAndServe(hosts[0], mux))
	}()

	listener, err := net.Listen(method, hosts[1])
	if err != nil {
		log.Fatalln(err.Error())
	}

	grpcServer := grpc.NewServer()
	notes.RegisterNotesServer(grpcServer, app.NewNotes())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err.Error())
	}
}
