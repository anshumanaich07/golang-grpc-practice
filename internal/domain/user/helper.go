package user

import (
	"learn-grpc/internal/api/grpc/pb"
	"strconv"

	"github.com/pkg/errors"
)

func ConvertUserIDReqToInt(id *pb.UserIDRequest) (int, error) {
	i, err := strconv.Atoi(id.Id)
	if err != nil {
		return 0, errors.Wrap(err, "unable to convert id to int")
	}

	return i, nil
}

func ConvertToUser(u *pb.User) (*User, error) {
	user := &User{}
	if u.Id == "" {
		return nil, errors.New("ID is empty")
	}
	uid, err := strconv.Atoi(u.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert string ID to int")
	}
	user.ID = uint(uid)
	user.Email = u.Email
	user.Name = u.Name

	return user, nil
}
