package main

import (
	"context"
	"fmt"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
)

const (
	maxReconnectAttempts = 10
	initialBackoff     = 1 * time.Second
	maxBackoff         = 5 * time.Second
)

func createGRPCClient(address string) (*grpc.ClientConn, error) {
	backoffConfig := backoff.Config{
		Initial: initialBackoff,
		Max:     maxBackoff,
	}
	connectParams := grpc.ConnectParams{
		Backoff: backoffConfig,
		MinConnectTimeout: 5 * time.Second,
	}
	clientConn, err := grpc.Dial(address, grpc.WithConnectParams(connectParams), grpc.WithBackoffConfig(backoffConfig))
	if err != nil {
		return nil, err
	}
	return clientConn, nil
}

func main() {
	address := "localhost:50051"
	clientConn, err := createGRPCClient(address)
	if err != nil {
		fmt.Println("Failed to create gRPC client:", err)
		return
	}
	defer clientConn.Close()
	fmt.Println("gRPC client created successfully")
}