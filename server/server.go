package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc-gateway/internal/handler"
	"grpc-gateway/pkg/gen/notes"
	"log"
	"net"
	"net/http"
)

const (
	method = "tcp"
)

type Server struct {
	socketOne string
	socketTwo string
}

func NewServer(socketOne, socketTwo string) *Server {

	return &Server{
		socketOne: socketOne,
		socketTwo: socketTwo,
	}
}

func (s Server) RunAPIServer() {

	go func() {

		mux := runtime.NewServeMux()
		notes.RegisterNotesHandlerServer(context.Background(), mux, handler.NewNotes())
		log.Fatalln(http.ListenAndServe(s.socketOne, mux))
	}()

	listener, err := net.Listen(method, s.socketTwo)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	notes.RegisterNotesServer(grpcServer, handler.NewNotes())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
