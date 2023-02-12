package user

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, request models.User) (*models.User, error)
	UpdateStatusUser(ctx context.Context, id, status string) error
	GetOneByEmail(ctx context.Context, email string) (*models.User, error)
	GetVerifyByEmail(ctx context.Context, email string) (*models.EmailVerification, error)
}
type user_repository struct {
	loggger *logrus.Logger
	gorm    *gorm.DB
}

func NewUserRepositorys(config *config.Config) *user_repository {
	return &user_repository{loggger: config.Logger, gorm: config.Gorm}
}
