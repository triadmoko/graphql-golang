package comment

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *comment_repository) Detail(ctx context.Context, id, userID string) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	comment := models.Comment{}
	err := r.gorm.First(&comment, "deleted_at IS NULL AND id = ? AND user_id = ?", id, userID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
