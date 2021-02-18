package server

import (
	"fmt"
	"log"
	"net"

	"github.com/erickmaria/kangaroo/core/pkg/logger"
	"github.com/erickmaria/kangaroo/core/pkg/profile"
	"google.golang.org/grpc"
)

type KangarooServer struct {
	module  string
	address string
	Grpc    *grpc.Server
}

func NewKangarooServer(address string) *KangarooServer {

	return &KangarooServer{
		address: address,
	}
}

func (svr *KangarooServer) Listen() *KangarooServer {

	fmt.Println("Initializing Application")

	fmt.Println("Starting GRPC Server...")
	lis, err := net.Listen("tcp", svr.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svr.Grpc = grpc.NewServer()

	fmt.Printf("Server listening on %s\n", svr.address)
	if err := svr.Grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return svr
}

func (svr *KangarooServer) SetProperties(properties interface{}, profilesPath string, envPrefix string) *KangarooServer {

	err := profile.Init(properties, profilesPath, envPrefix)
	if err != nil {
		log.Fatalf("failed to init properties: %v", err)
	}

	return svr
}

func (svr *KangarooServer) SetLoggerModule(name string) *KangarooServer {

	if len(name) == 0 {
		name = ""
	}

	logger.SetModule(name)

	return svr
}
