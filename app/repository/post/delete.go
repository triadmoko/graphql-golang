package post

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *post_repository) Delete(ctx context.Context, id, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()
	now := time.Now().UTC()

	err := r.gorm.WithContext(ctx).Model(&models.Post{}).Where("deleted_at IS NULL AND id = ? AND user_id = ?", id, userID).Updates(
		&models.Post{
			DeletedAt: &now,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}
