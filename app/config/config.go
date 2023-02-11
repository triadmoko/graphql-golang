package config

import (
	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/logs"
	"gorm.io/gorm"
)

type (
	Config struct {
		Logger *logrus.Logger
		Gorm   *gorm.DB
	}
)

func InitConfig() (*Config, error) {
	logger := logs.NewLogger("./logs/logger.log", true)
	gorm, err := InitPostgreGorm(logger)
	if err != nil {
		logger.Error("Error init gorm", err)
		return nil, err
	}

	response := Config{
		Logger: logger,
		Gorm:   gorm,
	}
	return &response, nil
}
