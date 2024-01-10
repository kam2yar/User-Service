package internal

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/kam2yar/user-service/api"
	v1 "github.com/kam2yar/user-service/internal/handlers/rpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"
)

var (
	httpPort = 80
	grpcPort = 8080
)

func Bootstrap() {
	go func() {
		if err := ServeGRPCGateway(); err != nil {
			grpclog.Fatal(err)
		}
	}()

	serveGRPC()
}

func serveGRPC() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &v1.UserManagementServer{})

	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func ServeGRPCGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		return err
	}

	log.Printf("grpc gateway server listening at %d", httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}
