package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServer) GetUserProfile(
	ctx context.Context,
	req *pb.UserRequest,
) (*pb.UserResponse, error) {
	fmt.Printf("Permintaan profil diterima untuk user ID: %d\n", req.GetUserId())

	if req.GetUserId() != 1 {
		return nil, status.Error(codes.NotFound, "User tidak ditemukan")
	}

	return &pb.UserResponse{
		Name:  "Budi Santoso",
		Email: "budi@universitas.ac.id",
		Phone: "081234567890",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userServer{})

	fmt.Println("Server gRPC berjalan pada port 50051")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
