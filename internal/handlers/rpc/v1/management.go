package v1

import (
	"context"
	"fmt"
	pb "github.com/kam2yar/user-service/api"
	"math/rand"
)

type UserManagementServer struct {
	pb.UnimplementedUserServer
}

func (s *UserManagementServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	fmt.Println("CTX:", ctx)
	fmt.Println("Request:", request)
	var userId = int32(rand.Intn(10000))

	return &pb.CreateResponse{Id: userId, Name: request.Name, Email: request.Email, Password: request.Password, CreatedAt: "now"}, nil
}
