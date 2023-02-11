package user

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *user_repository) Create(ctx context.Context, request models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.User{}).Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
