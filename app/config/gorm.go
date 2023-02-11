package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreGorm(logger *logrus.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		helpers.LoadEnv("DB_HOST"),
		helpers.LoadEnv("DB_USER"),
		helpers.LoadEnv("DB_PASS"),
		helpers.LoadEnv("DB_NAME"),
		helpers.LoadEnv("DB_PORT"),
		helpers.LoadEnv("SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("not connect gorm postgre", err.Error())
		return nil, err
	}
	return db, nil
}
