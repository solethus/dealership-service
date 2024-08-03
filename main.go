package main

import (
	"database/sql"

	"net"

	"google.golang.org/grpc"

	"github.com/solethus/dealership-service/internal/respository"
	"github.com/solethus/dealership-service/internal/server"
	"github.com/solethus/dealership-service/internal/service"
	log "github.com/solethus/dealership-service/pkg/logger"
	pb "github.com/solethus/dealership-service/proto/dealership"
)

func main() {

	// Set up database connection
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dealership?sslmode=disable")
	if err != nil {
		log.Logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	// Create repository, service, and server

	repo := respository.NewDealershipRepository(db)
	svc := service.NewDealershipService(repo)
	srv := server.NewDealershipServer(svc)

	//// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterDealershipServiceServer(grpcServer, srv)

	//// Start listening
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Logger.Fatalf("Failed to listen: %v", err)
	}

	log.Logger.Info("Server listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Logger.Fatalf("Failed to serve: %v", err)
	}
}
