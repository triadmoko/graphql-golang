package comment

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *comment_repository) Delete(ctx context.Context, id, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()
	now := time.Now().UTC()

	err := r.gorm.WithContext(ctx).Model(&models.Comment{}).Where("deleted_at IS NULL AND id = ? AND user_id = ?", id, userID).Updates(
		&models.Comment{
			DeletedAt: &now,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}
