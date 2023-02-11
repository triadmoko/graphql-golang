package user

import (
	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/config"
	"gorm.io/gorm"
)

type UserRepository interface {
}
type user_repository struct {
	loggger *logrus.Logger
	gorm    *gorm.DB
}

func NewUserRepositorys(config *config.Config) *user_repository {
	return &user_repository{loggger: config.Logger, gorm: config.Gorm}
}
