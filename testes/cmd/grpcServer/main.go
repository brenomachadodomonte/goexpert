package main

import (
	"go-expert/testes/internal/service"
	"go-expert/testes/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	distanceService := service.NewDistanceService()

	grpcServer := grpc.NewServer()
	pb.RegisterDistanceServiceServer(grpcServer, distanceService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	log.Println("gRPC Server listening on port 50051")
	log.Fatal(grpcServer.Serve(lis))
}
