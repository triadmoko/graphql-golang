package comment

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *comment_repository) Update(ctx context.Context, id string, comment models.Comment) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.Comment{}).Where("deleted_at IS NULL AND id = ?", id).Updates(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
