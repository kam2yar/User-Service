package internal

import (
	"fmt"
	pb "github.com/kam2yar/user-service/api/pb"
	v1 "github.com/kam2yar/user-service/internal/handlers/rpc/v1"
	"google.golang.org/grpc"
	"log"
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
	pb.RegisterUserServer(s, &v1.UserManagementServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
