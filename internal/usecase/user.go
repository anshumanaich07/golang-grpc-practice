package usecase

import (
	"context"
	"fmt"
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
	fmt.Println("user usecase")
	userUC.UserRepo.GetUserByID(ctx, 0)
	return nil, nil
}
