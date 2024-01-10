package v1

import (
	"context"
	pb "github.com/kam2yar/user-service/api"
	"github.com/kam2yar/user-service/internal/dto"
	"github.com/kam2yar/user-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserManagementServer struct {
	pb.UnimplementedUserServer
}

func (s *UserManagementServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.UserData, error) {
	userDto := dto.UserDto{}
	userDto.SetName(request.GetName())
	userDto.SetEmail(request.GetEmail())
	userDto.SetPassword(request.GetPassword())

	err := services.CreateUser(&userDto)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.UserData{
		Id:        uint32(userDto.GetId()),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: userDto.GetUpdatedAt().Format(time.DateTime),
	}, nil
}

func (s *UserManagementServer) Find(ctx context.Context, request *pb.FindRequest) (*pb.UserData, error) {
	userDto, err := services.FindUser(uint(request.GetId()))
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.UserData{
		Id:        uint32(userDto.GetId()),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: userDto.GetUpdatedAt().Format(time.DateTime),
	}, nil
}

func (s *UserManagementServer) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	userList := services.List(int(request.GetLimit()))

	var users = make([]*pb.UserData, len(*userList))
	for i, userDto := range *userList {
		users[i] = &pb.UserData{
			Id:        uint32(userDto.GetId()),
			Name:      userDto.GetName(),
			Email:     userDto.GetEmail(),
			CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
			UpdatedAt: userDto.GetUpdatedAt().Format(time.DateTime),
		}
	}

	return &pb.ListResponse{
		Users: users,
	}, nil
}

func (s *UserManagementServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UserData, error) {
	userDto := dto.UserDto{}
	userDto.SetId(uint(request.GetId()))
	userDto.SetName(request.GetName())
	userDto.SetEmail(request.GetEmail())
	userDto.SetPassword(request.GetPassword())

	err := services.UpdateUser(&userDto)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.UserData{
		Id:        uint32(userDto.GetId()),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		CreatedAt: userDto.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: userDto.GetUpdatedAt().Format(time.DateTime),
	}, nil
}

func (s *UserManagementServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := services.DeleteUser(uint(request.GetId()))

	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.DeleteResponse{
		Success: true,
	}, nil
}
