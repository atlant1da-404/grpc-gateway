package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc-gateway/config"
	"grpc-gateway/grpc/gen/notes"
	"grpc-gateway/pkg/handler"
	"log"
	"net"
	"net/http"
)

const (
	method = "tcp"
)

func main() {

	socketOne, socketTwo := config.LoadSocketConfig()
	if socketOne == "" && socketTwo == "" {
		socketOne = "localhost:8080"
		socketTwo = "localhost:8081"
	}

	go func() {

		mux := runtime.NewServeMux()
		notes.RegisterNotesHandlerServer(context.Background(), mux, handler.NewNotes())
		log.Fatalln(http.ListenAndServe(socketOne, mux))
	}()

	listener, err := net.Listen(method, socketTwo)
	if err != nil {
		log.Fatalln(err.Error())
	}

	grpcServer := grpc.NewServer()
	notes.RegisterNotesServer(grpcServer, handler.NewNotes())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err.Error())
	}
}
