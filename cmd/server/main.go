package main

import (
	"fmt"
	"grpc/internal/service"
	pb "grpc/pb/math"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":9100")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen on :9100: %v\n", err)
		os.Exit(1)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMathUtilServer(grpcServer, &service.UtilServer{})
	fmt.Printf("RankSystem gRPC server listening on %s\n", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Fprintf(os.Stderr, "failed to serve: %v\n", err)
		os.Exit(1)
	}
}
