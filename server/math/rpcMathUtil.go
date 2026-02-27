package math

import (
	"context"
)

type UtilServer struct {
	UnimplementedMathUtilServer
}

func (s *UtilServer) Add(ctx context.Context, req *Vector2) (*Num, error) {
	return &Num{Num: req.Numa + req.Numb}, nil
}

func (s *UtilServer) Sub(ctx context.Context, req *Vector2) (*Num, error) {
	return &Num{Num: req.Numa - req.Numb}, nil
}
