package main

import (
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port_grpc = ":51234"
	port_rest = ":5123"
	n_jobs    = 10
)

func main() {
	log.Println("Starting in-memory storage")
	start_grpc()
	start_rest()
}

func start_grpc() {
	listener, err := net.Listen("tcp", port_grpc)
	if err != nil {
		log.Panicln("Error listening on GRPC port", port_grpc, ":", err)
	}
	s := grpc.NewServer()
	api.RegisterBotDBServer(s, db.NewServer(n_jobs))
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Panicln("Error serving GRPC", err)
	}
}

func start_rest() {
	mux := runtime.NewServeMux()
	err := api.RegisterBotDBHandlerFromEndpoint(context.Background(), mux, port_grpc, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Panicln("Error connecting gateway to GRPC", err)
	}
	server := http.Server{
		Handler: mux,
	}
	l, err := net.Listen("tcp", port_rest)
	if err != nil {
		log.Panicln("Error listening on REST port", port_rest, err)
	}
	err = server.Serve(l)
	if err != nil {
		log.Panicln("Error serving REST")
	}
}
