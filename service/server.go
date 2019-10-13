package service

import (
	"context"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/brunoluiz/snapurl"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func StartGRPCservice(address string) error {
	// create a listener on TCP port 5000
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a service instance
	h := Service{}

	// create a gRPC service object
	server := grpc.NewServer()
	snapurl.RegisterURLSnapServer(server, &h)

	log.Infof("GRPC service on %s", address)
	return server.Serve(lis)
}

func StartGRPCGateway(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	runtime.SetHTTPBodyMarshaler(mux)

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping
	err := snapurl.RegisterURLSnapHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("could not register service Ping: %s", err)
	}

	log.Infof("GRPC Gateway on %s (proxy to %s)", address, grpcAddress)
	return http.ListenAndServe(address, mux)
}
