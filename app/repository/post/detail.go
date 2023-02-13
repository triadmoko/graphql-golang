package post

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *post_repository) Detail(ctx context.Context, id, userID string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	post := models.Post{}
	err := r.gorm.First(&post, "deleted_at IS NULL AND id = ? AND user_id = ?", id, userID).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}
