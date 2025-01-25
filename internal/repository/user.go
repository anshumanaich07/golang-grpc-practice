package repository

import (
	"context"
	"fmt"
	"learn-grpc/internal/domain/user"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) (*userRepo, error) {
	return &userRepo{
		db: db,
	}, nil
}

func (ur *userRepo) GetUserByID(ctx context.Context, userID int) (*user.User, error) {
	fmt.Println("user repo")
	return nil, nil
}
