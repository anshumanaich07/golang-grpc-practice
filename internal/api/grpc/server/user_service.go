package grpc

import (
	"context"
	"learn-grpc/internal/api/grpc/pb"
	"learn-grpc/internal/domain/user"
	"strconv"

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
	u, err := usi.UserUsecase.GetUserByID(ctx, idInt)
	if err != nil || u == nil {
		return nil, errors.Wrap(err, "unable to find user")
	}

	convUser := user.ConvertUserToPbUser(*u)

	return convUser, nil
}

func (usi *UserServiceImpl) GetAllUsers(ctx context.Context, empty *pb.EmptyRequest) (*pb.AllUsersResponse, error) {
	users, err := usi.UserUsecase.GetUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find users")
	}

	response := pb.AllUsersResponse{}
	for _, u := range users {
		response.Users = append(response.Users, user.ConvertUserToPbUser(*u))
	}

	return &response, nil
}

func (usi *UserServiceImpl) UpdateUser(ctx context.Context, updUserReq *pb.User) (*pb.User, error) {
	if updUserReq.Email == "" || updUserReq.Name == "" || updUserReq.Id == "" {
		return nil, errors.New("name, email or id cannot be empty")
	}

	// convert to model user
	temp := pb.AddUserReq{}
	temp.Email = updUserReq.Email
	temp.Name = updUserReq.Name
	userModel, err := user.ConvertToUser(&temp)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert to User")
	}

	idInt, err := strconv.Atoi(updUserReq.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert id from string to int")
	}

	updatedUser, err := usi.UserUsecase.UpdateUser(ctx, idInt, *userModel)
	if err != nil {
		return nil, errors.Wrap(err, "unable to updated user")
	}

	convUser := user.ConvertUserToPbUser(*updatedUser)

	return convUser, nil
}

func (usi *UserServiceImpl) DeleteUser(ctx context.Context, delReq *pb.UserIDRequest) (*pb.DeleteUserResponse, error) {
	if delReq.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	// convert id to int
	idInt, err := strconv.Atoi(delReq.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert id from string to int")
	}
	err = usi.UserUsecase.DeleteUser(ctx, idInt)
	if err != nil {
		return nil, errors.Wrap(err, "unable to delete user")
	}

	return &pb.DeleteUserResponse{Id: delReq.Id}, nil
}
