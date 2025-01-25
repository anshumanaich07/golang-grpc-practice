package user

import "context"

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:200;not null"`
	Email string `gorm:"size:100;not null"`
}

type UserUsecase interface {
	GetUserByID(ctx context.Context, userID int) (*User, error)
}

type UserRepository interface {
	GetUserByID(ctx context.Context, userID int) (*User, error)
}
