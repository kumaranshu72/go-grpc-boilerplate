package main

import (
	"context"
	"flag"
	v1 "live-tracking/pkg/api/v1"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req1 := v1.EmptyRequest{
		Api: apiVersion,
	}

	res1, err := c.GetHealth(ctx, &req1)
	if err != nil {
		log.Fatalf("HealthCheck failed: %v", err)
	}
	log.Printf("Health Check Result: <%+v>\n\n", res1)
}
