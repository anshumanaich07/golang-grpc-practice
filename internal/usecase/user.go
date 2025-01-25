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
	userUC.UserRepo.GetUserByID(ctx, 0)
	return nil, nil
}

func (userUC *userUsecase) AddUser(ctx context.Context, user user.User) (*user.User, error) {
	return userUC.UserRepo.AddUser(ctx, user)
}
