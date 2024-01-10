package v1

import (
	"context"
	pb "github.com/kam2yar/user-service/api/pb"
	"github.com/kam2yar/user-service/internal/dto"
	"github.com/kam2yar/user-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserManagementServer struct {
	pb.UnimplementedUserServer
}

func (s *UserManagementServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	userDto := dto.UserDto{}
	userDto.SetName(request.GetName())
	userDto.SetEmail(request.GetEmail())
	userDto.SetPassword(request.GetPassword())

	err := services.CreateUser(&userDto)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.CreateResponse{
		Id:        uint32(userDto.GetId()),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
	}, nil
}

func (s *UserManagementServer) Find(ctx context.Context, request *pb.FindRequest) (*pb.FindResponse, error) {
	userDto, err := services.FindUser(uint(request.GetId()))
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.FindResponse{
		Id:        uint32(userDto.GetId()),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: userDto.GetUpdatedAt().Format(time.DateTime),
	}, nil
}
