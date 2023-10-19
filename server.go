package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"stocks_service/stocks"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on 9000: %v", err)
	}

	s := stocks.Server{}
	grpcServer := grpc.NewServer()

	stocks.RegisterStocksServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to server on port 9000: %v", err)
	}

}
