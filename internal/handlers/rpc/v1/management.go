package v1

import (
	"context"
	pb "github.com/kam2yar/user-service/api"
	"github.com/kam2yar/user-service/internal/dto"
	"github.com/kam2yar/user-service/internal/services"
	"time"
)

type UserManagementServer struct {
	pb.UnimplementedUserServer
}

func (s *UserManagementServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	userDto := dto.UserDto{}
	userDto.SetName(request.Name)
	userDto.SetEmail(request.Email)
	userDto.SetPassword(request.Password)

	services.CreateUser(&userDto)

	return &pb.CreateResponse{
		Id:        uint32(userDto.GetId()),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
	}, nil
}
