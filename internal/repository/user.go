package repository

import (
	"context"
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
	var user user.User
	result := ur.db.First(&user, userID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "unable to find user by ID")
	}

	return &user, nil
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

func (ur *userRepo) GetUsers(ctx context.Context) ([]*user.User, error) {
	users := []*user.User{}
	result := ur.db.Find(&users)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "unable to find the users")
	}

	return users, nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, id int, updUser user.User) (*user.User, error) {
	result := ur.db.Model(&user.User{}).Where("id = ?", id).Updates(updUser)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "unable to update user")
	}

	// get updated user
	updatedUser, err := ur.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(result.Error, "unable to find the user based on id")
	}

	return updatedUser, nil
}

func (ur *userRepo) DeleteUser(ctx context.Context, id int) error {
	result := ur.db.Delete(&user.User{}, id)
	if result.Error != nil {
		return errors.Wrap(result.Error, "unable to delete user")
	}

	return nil
}
