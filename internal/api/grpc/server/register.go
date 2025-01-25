package grpc

import (
	"fmt"
	"learn-grpc/internal/config"
	"learn-grpc/internal/repository"
	"learn-grpc/internal/usecase"
	"log"
	"net"

	pb "learn-grpc/internal/api/grpc/pb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func RegisterUserService(grpcConfig *config.GRPCConfig, db *gorm.DB) error {
	lis, err := net.Listen(grpcConfig.Protocol, fmt.Sprintf(":%d", grpcConfig.Port))
	if err != nil {
		return errors.Wrap(err, "unable to listen on port")
	}

	// grpc server
	grpcServer := grpc.NewServer()

	// register methods
	userRepo, err := repository.NewUserRepo(db)
	if err != nil {
		return errors.Wrap(err, "unable to get gorm DB object")
	}
	userUC, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		return errors.Wrap(err, "unable to get user usecase")
	}
	userSrv := NewUserServiceImpl(userUC)

	pb.RegisterUserServiceServer(grpcServer, userSrv)

	// reflection: exposes the methods available
	reflection.Register(grpcServer)

	log.Printf("gRPC server is running on port %d\n", grpcConfig.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server %v", err)
	}

	return nil
}
