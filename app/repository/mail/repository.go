package mail

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

type MailRepository interface {
	Create(ctx context.Context, request models.EmailVerification) (*models.EmailVerification, error)
}
type mail_repository struct {
	loggger *logrus.Logger
	gorm    *gorm.DB
}

func NewMailRepositorys(config *config.Config) *mail_repository {
	return &mail_repository{loggger: config.Logger, gorm: config.Gorm}
}
