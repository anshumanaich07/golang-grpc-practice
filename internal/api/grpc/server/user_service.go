package grpc

import (
	"context"
	"learn-grpc/internal/api/grpc/pb"
	"learn-grpc/internal/domain/user"

	"github.com/pkg/errors"
)

// this struct will implement all the interface methods in 'user_grpc.pb.go'
type UserServiceImpl struct {
	UserUsecase user.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserServiceImpl(userUsecase user.UserUsecase) *UserServiceImpl {
	return &UserServiceImpl{
		UserUsecase: userUsecase,
	}
}

func (usi *UserServiceImpl) AddUser(ctx context.Context, userReq *pb.AddUserReq) (*pb.User, error) {
	// validate userReq
	if userReq.Email == "" || userReq.Name == "" {
		return nil, errors.New("Email or Name cannot be empty")
	}

	userModel, err := user.ConvertToUser(userReq)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert to User")
	}

	respUser, err := usi.UserUsecase.AddUser(ctx, *userModel)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert to User")
	}
	convUser := user.ConvertUserToPbUser(*respUser)

	// convert to pb User
	return convUser, nil
}
func (usi *UserServiceImpl) GetUser(ctx context.Context, userID *pb.UserIDRequest) (*pb.User, error) {
	// validate the request
	if userID.Id == "" {
		return nil, errors.New("user ID request cannot be empty")
	}
	idInt, err := user.ConvertUserIDReqToInt(userID)
	if err != nil {
		return nil, err
	}
	usi.UserUsecase.GetUserByID(ctx, idInt)

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
