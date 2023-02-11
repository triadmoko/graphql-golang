package user

import (
	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/model"
	"github.com/triadmoko/grahpql-golang/repository/user"
)

type UserServices interface {
	Create(request model.NewUser) (model.User, error)
}
type user_service struct {
	loggger        *logrus.Logger
	userRepository user.UserRepository
}

func NewUserServices(loggger *logrus.Logger, userRepository user.UserRepository) *user_service {
	return &user_service{loggger: loggger, userRepository: userRepository}
}
