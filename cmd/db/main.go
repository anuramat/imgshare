package main

import (
	"log"
	"net"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":51234"
)

func main() {
	log.Println("Starting in-memory storage")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicln("Error starting @ port", port, ":", err)
	}
	s := grpc.NewServer()
	api.RegisterBotDBServer(s, db.NewServer())
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Panicln("Error serving", err)
	}
}
