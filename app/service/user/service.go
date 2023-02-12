package user

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/model"
	"github.com/triadmoko/grahpql-golang/repository/user"
	"github.com/triadmoko/grahpql-golang/service/mail"
)

type UserServices interface {
	Create(ctx context.Context, req model.NewUser) (*model.User, error)
	Login(ctx context.Context, req model.NewLogin) (*model.Token, error)
}
type user_service struct {
	loggger        *logrus.Logger
	userRepository user.UserRepository
	mailService    mail.MailServices
}

func NewUserServices(
	loggger *logrus.Logger,
	userRepository user.UserRepository,
	mailService mail.MailServices,
) *user_service {
	return &user_service{loggger: loggger, userRepository: userRepository, mailService: mailService}
}
