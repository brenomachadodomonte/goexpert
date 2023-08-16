package main

import (
	"database/sql"
	"github.com/brenomachadodomonte/go-expert/grpc/internal/database"
	"github.com/brenomachadodomonte/go-expert/grpc/internal/pb"
	"github.com/brenomachadodomonte/go-expert/grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	log.Println("gRPC Server running on port 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
