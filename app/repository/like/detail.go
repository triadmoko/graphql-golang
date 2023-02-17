package like

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *like_repository) Detail(ctx context.Context, userID, postID string) (*models.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	like := models.Like{}
	err := r.gorm.First(&like, "deleted_at IS NULL AND user_id = ? AND post_id = ?", userID, postID).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}
