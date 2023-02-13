package post

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *post_repository) Update(ctx context.Context, id string, post models.Post) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.Post{}).Where("deleted_at IS NULL AND id = ?", id).Updates(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}
