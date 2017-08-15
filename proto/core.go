package push

import (
	"flag"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
)

type Server struct {
	Port *int
	App *string
}

func NewServer() *Server {
	return &Server {
		Port: flag.Int("port", 8000, "The server port"),
		App: flag.String("app", "default", "The application identifier"),
	}
}

func (s *Server)Run(f func(*grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *s.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("listen %v port", *s.Port)
	}

	opts := []grpc.ServerOption{}
	ns := grpc.NewServer(opts...)

	f(ns)

	log.Println("start gRPC...")
	ns.Serve(lis)
}
