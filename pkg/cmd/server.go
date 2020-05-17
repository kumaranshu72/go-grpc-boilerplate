package cmd

import (
	"context"
	"fmt"
	"live-tracking/pkg/logger"
	"live-tracking/config"
	"live-tracking/pkg/protocol/grpc"
	"live-tracking/pkg/protocol/rest"
	v1 "live-tracking/pkg/service/v1"
)

const (
	configPath = "config/reader"
)
// RunServer runs grpc server and Http gateway
func RunServer() error {
	ctx := context.Background()

	config.GetConfig(configPath)

	// initialize logger
	if err := logger.Init(config.Config.LogLevel, config.Config.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	v1Api := v1.NewHealthService()

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, config.Config.GRPCPort, config.Config.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1Api, config.Config.GRPCPort)
}
