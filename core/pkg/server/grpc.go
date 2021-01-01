package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	address string
	Grpc    *grpc.Server
}

func NewServer(address string) *Server {

	return &Server{
		address: address,
	}
}

func (svr *Server) Listen() *Server {

	fmt.Println("Starting GRPC Server")
	lis, err := net.Listen("tcp", svr.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svr.Grpc = grpc.NewServer()

	fmt.Printf("Serve on %s\n", svr.address)
	if err := svr.Grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return svr
}
