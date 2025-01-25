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

func ConvertToUser(u *pb.AddUserReq) (*User, error) {
	user := &User{}
	user.Email = u.GetEmail()
	user.Name = u.GetName()

	return user, nil
}

func ConvertUserToPbUser(u User) *pb.User {
	user := &pb.User{}
	user.Id = strconv.Itoa(int(u.ID))
	user.Email = u.Email
	user.Name = u.Name

	return user
}
