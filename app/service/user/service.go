package user

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/repository/user"
	"github.com/triadmoko/grahpql-golang/service/mail"
)

type UserServices interface {
	Create(ctx context.Context, req request.NewUser) (*request.User, error)
	Detail(ctx context.Context, id string) (*request.User, error)
	Update(ctx context.Context, id string, req *request.UpdateUser) (*request.User, error)

	Login(ctx context.Context, req request.NewLogin) (*request.Token, error)
	VerifyEmail(ctx context.Context, req request.NewVerify) (string, error)
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
