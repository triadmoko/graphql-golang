package config

import (
	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type (
	Config struct {
		Logger *logrus.Logger
		Gorm   *gorm.DB
		Mongo  *mongo.Database
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
		Mongo:  InitMongoDB(),
	}
	return &response, nil
}
