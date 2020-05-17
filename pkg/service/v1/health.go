package v1

import (
	"context"
	v1 "live-tracking/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// healthServiceServer is implementation of v1.healthServiceServer proto interface
type healthServiceServer struct {
}

// NewHealthService creates new health service
func NewHealthService() {
	return &healthServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *healthServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented, "unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// GetHealth checks the health of the server
func (s *healthServiceServer) GetHealth(ctx context.Context, req *v1.EmptyRequest) (*v1.HealthResponse, error) {
	data := &v1.HealthResponse_Data{
		Message: "server is live",
	}
	return &v1.HealthResponse{
		StatusCode:    codes.OK,
		StatusMessage: "success",
		Data:          data,
	}, nil
}
