package service

import (
	"context"
	"go-expert/testes/pkg/pb"
)

type DistanceService struct {
	pb.UnimplementedDistanceServiceServer
}

func NewDistanceService() *DistanceService {
	return &DistanceService{}
}

func (ds *DistanceService) GetDistance(ctx context.Context, in *pb.GetDistanceInput) (*pb.GetDistanceOutput, error) {
	return &pb.GetDistanceOutput{
		Distance: 10,
	}, nil
}
