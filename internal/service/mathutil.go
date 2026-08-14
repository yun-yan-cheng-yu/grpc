package service

import (
	"context"

	pb "grpc/pb/math"
)

// UtilServer implements the MathUtil gRPC service (handwritten business logic).
type UtilServer struct {
	pb.UnimplementedMathUtilServer
}

func (s *UtilServer) Add(ctx context.Context, req *pb.Vector2) (*pb.Num, error) {
	return &pb.Num{Num: req.Numa + req.Numb}, nil
}

func (s *UtilServer) Sub(ctx context.Context, req *pb.Vector2) (*pb.Num, error) {
	return &pb.Num{Num: req.Numa - req.Numb}, nil
}
