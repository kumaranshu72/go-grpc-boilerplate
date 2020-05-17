package cmd

import (
	"context"
	"flag"
	"fmt"
	"live-tracking/pkg/protocol/grpc"
	v1 "live-tracking/pkg/service/v1"
)

// Config is configuration for server
type Config struct {
	// gRPC Port String
	GRPCPort string
}

// RunServer runs grpc server and Http gateway
func RunServer() error {
	ctx := context.Background()

	// get Configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	v1Api := v1.NewHealthService()

	return grpc.RunServer(ctx, v1Api, cfg.GRPCPort)
}
