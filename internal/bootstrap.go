package internal

import (
	"context"
	"fmt"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/kam2yar/user-service/api"
	v1 "github.com/kam2yar/user-service/internal/handlers/rpc/v1"
	"github.com/kam2yar/user-service/internal/interceptors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
)

var (
	httpPort        = 80
	grpcPort        = 8080
	recoveryHandler grpcRecovery.RecoveryHandlerFunc
)

func Bootstrap() {
	go func() {
		if err := ServeGRPCGateway(); err != nil {
			grpclog.Fatal(err)
		}
	}()

	ServeGRPC()
}

func ServeGRPC() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := zap.NewExample()
	recoveryHandler = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	recoveryOpts := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(recoveryHandler),
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptors.LoggerInterceptor(logger), logging.WithLogOnEvents(logging.StartCall, logging.FinishCall)), grpcRecovery.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(recoveryOpts...),
		))
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
