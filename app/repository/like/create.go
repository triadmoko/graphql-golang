package like

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *like_repository) Create(ctx context.Context, request models.Like) (*models.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.Like{}).Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
