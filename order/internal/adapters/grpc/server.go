package grpc

import (
	"fmt"
	"log"
	"net"
	"order/config"
	"order/internal/ports"

	"github.com/dctrseltsam/microservices-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	order.UnimplementedOrderServer
}

func New(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %v, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer(
	//grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	a.server = grpcServer
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting server on port %v ...", a.port)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %v", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
