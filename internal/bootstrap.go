package internal

import (
	"context"
	"fmt"
	pb "github.com/kam2yar/user-service/api"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

func Bootstrap() {
	serveGRPC()
}

func serveGRPC() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserServer
}

func (s *UserManagementServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	fmt.Println("CTX:", ctx)
	fmt.Println("Request:", request)
	var user_id int32 = int32(rand.Intn(10000))

	return &pb.CreateResponse{Id: user_id, Name: request.Name, Email: request.Email, CreatedAt: "now"}, nil
}
