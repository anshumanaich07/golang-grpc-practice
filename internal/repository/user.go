package repository

import (
	"context"
	"fmt"
	"learn-grpc/internal/domain/user"

	"github.com/pkg/errors"
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

func (ur *userRepo) AddUser(ctx context.Context, userReq user.User) (*user.User, error) {
	result := ur.db.Create(&userReq)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "unable to insert User")
	}

	// get from the table via id
	u := &user.User{}
	result = ur.db.First(&u, userReq.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "unable to find the user based on id")
	}

	return u, nil
}
