package grpc

import (
	"context"
	"fmt"
	"learn-grpc/internal/api/grpc/pb"
)

// this struct will implement all the interface methods in 'user_grpc.pb.go'
type UserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

func (usi *UserServiceImpl) GetUser(ctx context.Context, userID *pb.UserIDRequest) (*pb.User, error) {
	fmt.Println("reached the grpc method 'GetUser'")
	fmt.Println("ReqID: ", userID)
	// res := &pb.User{
	// 	Id:    userID.Id,
	// 	Name:  "Anshuman Aich",
	// 	Email: "anshumanaich@gmail.com",
	// }

	// convert the request
	// validate the request
	// call the usecase

	return nil, nil
}

func (usi *UserServiceImpl) GetAllUsers(ctx context.Context, empty *pb.EmptyRequest) (*pb.AllUsersResponse, error) {
	return nil, nil
}

func (usi *UserServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return nil, nil
}

func (usi *UserServiceImpl) UpdateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	return nil, nil
}

func (usi *UserServiceImpl) DeleteUser(ctx context.Context, delReq *pb.UserIDRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}
