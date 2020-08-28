package main

import (
	"github.com/jsanda/user-svc/pkg/user"
	"github.com/jsanda/user-svc/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	log.Println("starting")
	grpcPort := os.Getenv("GRPC_PORT")

	listener, err := net.Listen("tcp", ":" + grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	log.Println("creating UserService")
	userSvc, err := user.NewUserService(os.Getenv("CASSANDRA_SVC"))
	if err != nil {
		log.Fatalf("failed to create UserService: %s", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userSvc)
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	server.Serve(listener)
}
