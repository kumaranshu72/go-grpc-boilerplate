package cmd

import (
	"context"
	"flag"
	"fmt"
	"live-tracking/pkg/protocol/grpc"
	"live-tracking/pkg/protocol/rest"
	v1 "live-tracking/pkg/service/v1"
)

// Config is configuration for server
type Config struct {
	// gRPC Port String
	GRPCPort string
	HTTPPort string
}

// RunServer runs grpc server and Http gateway
func RunServer() error {
	ctx := context.Background()

	// get Configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	v1Api := v1.NewHealthService()

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1Api, cfg.GRPCPort)
}
