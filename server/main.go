package main

import (
	"fmt"
	"grpc/server/math"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, _ := net.Listen("tcp", ":9090")
	grpcServer := grpc.NewServer()
	math.RegisterMathUtilServer(grpcServer, &math.UtilServer{})
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
