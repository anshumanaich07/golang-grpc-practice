package grpc

import (
	"fmt"
	"learn-grpc/internal/config"
	"log"
	"net"

	pb "learn-grpc/internal/api/grpc/pb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterService(grpcConfig *config.GRPCConfig) error {
	lis, err := net.Listen(grpcConfig.Protocol, fmt.Sprintf(":%d", grpcConfig.Port))
	if err != nil {
		return errors.Wrap(err, "unable to listen on port")
	}

	// grpc server
	grpcServer := grpc.NewServer()

	// register methods

	pb.RegisterUserServiceServer(grpcServer, &UserServiceImpl{})

	// reflection: exposes the methods available
	reflection.Register(grpcServer)

	log.Printf("gRPC server is running on port %d\n", grpcConfig.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server %v", err)
	}

	return nil
}
