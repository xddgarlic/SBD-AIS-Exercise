package main

import (
	"log"
	"time"

	"exc8/client"
	"exc8/server"
)

func main() {
	// -----------------------------------
	// Start gRPC server in a goroutine
	// -----------------------------------
	go func() {
		if err := server.StartGrpcServer(); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// Give the server time to start listening
	time.Sleep(1 * time.Second)

	// -----------------------------------
	// Start gRPC client
	// -----------------------------------
	grpcClient, err := client.NewGrpcClient()
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}

	if err := grpcClient.Run(); err != nil {
		log.Fatalf("client error: %v", err)
	}
}
