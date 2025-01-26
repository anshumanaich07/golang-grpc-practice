package user

import "context"

type User struct {
	ID    uint   `gorm:"primaryKey,autoIncrement"`
	Name  string `json:"name" gorm:"size:200;not null"`
	Email string `json:"email" gorm:"size:100;unique;not null"`
}

type UserUsecase interface {
	AddUser(ctx context.Context, user User) (*User, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
	GetUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id int, user User) (*User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserRepository interface {
	AddUser(ctx context.Context, user User) (*User, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
	GetUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id int, user User) (*User, error)
	DeleteUser(ctx context.Context, id int) error
}
