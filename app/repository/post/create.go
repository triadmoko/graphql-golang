package post

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *post_repository) Create(ctx context.Context, request models.Post) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.Post{}).Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
