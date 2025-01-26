package usecase

import (
	"context"
	"learn-grpc/internal/domain/user"
)

type userUsecase struct {
	UserRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) (*userUsecase, error) {
	return &userUsecase{
		UserRepo: userRepo,
	}, nil
}

func (userUC *userUsecase) GetUserByID(ctx context.Context, userID int) (*user.User, error) {
	return userUC.UserRepo.GetUserByID(ctx, userID)
}

func (userUC *userUsecase) AddUser(ctx context.Context, user user.User) (*user.User, error) {
	return userUC.UserRepo.AddUser(ctx, user)
}

func (userUC *userUsecase) GetUsers(ctx context.Context) ([]*user.User, error) {
	return userUC.UserRepo.GetUsers(ctx)
}

func (userUC *userUsecase) UpdateUser(ctx context.Context, id int, updUser user.User) (*user.User, error) {
	return userUC.UserRepo.UpdateUser(ctx, id, updUser)
}

func (userUC *userUsecase) DeleteUser(ctx context.Context, id int) error {
	return userUC.UserRepo.DeleteUser(ctx, id)
}
