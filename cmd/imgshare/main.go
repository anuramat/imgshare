package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/imgshare"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	port_grpc    = ":51234"
	port_rest    = ":5123"
	port_swagger = ":8080"
	n_jobs       = 10
	rpc_timeout  = 500 * time.Millisecond
)

func main() {
	log.Println("Starting in-memory storage")
	go start_rest()
	go start_grpc()
	log.Panicln(http.ListenAndServe(port_swagger, http.FileServer(http.Dir("api/"))))
}

func start_grpc() {
	listener, err := net.Listen("tcp", port_grpc)
	if err != nil {
		log.Panicln("Error listening on GRPC port", port_grpc, ":", err)
	}
	interceptor := grpc.UnaryInterceptor(timeoutInterceptor)
	s := grpc.NewServer(interceptor)
	api.RegisterBotDBServer(s, imgshare.NewServer(n_jobs))
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Panicln("Error serving GRPC", err)
	}
}

func start_rest() {
	mux := runtime.NewServeMux()
	dialoption := grpc.WithTransportCredentials(insecure.NewCredentials())
	err := api.RegisterBotDBHandlerFromEndpoint(context.Background(), mux, port_grpc, []grpc.DialOption{dialoption})
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

func timeoutInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, rpc_timeout)
	defer cancel()
	resp, err = handler(ctx, req)
	return
}
