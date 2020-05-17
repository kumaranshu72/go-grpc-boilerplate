package grpc

import (
	"context"
	v1 "live-tracking/pkg/api/v1"
	"live-tracking/pkg/logger"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// RunServer runs gRPC service to publish GetHealth service
func RunServer(ctx context.Context, v1Api v1.HealthServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	// adding ssl
	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}
	server := grpc.NewServer(opts...)
	v1.RegisterHealthServer(server, v1Api)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
