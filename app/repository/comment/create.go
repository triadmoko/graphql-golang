package comment

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *comment_repository) Create(ctx context.Context, request models.Comment) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.Comment{}).Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
