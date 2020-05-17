package grpc

import (
	"context"
	v1 "live-tracking/pkg/api/v1"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish GetHealth service
func RunServer(ctx context.Context, v1Api v1.HealthServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterHealthServer(server, v1Api)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting Down GRPC Server")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("Starting gRPC Server...")
	return server.Serve(listen)
}