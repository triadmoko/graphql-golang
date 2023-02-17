package like

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *like_repository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()
	now := time.Now().UTC()

	err := r.gorm.WithContext(ctx).Model(&models.Like{}).Where("deleted_at IS NULL AND id = ? ", id).Updates(
		&models.Like{
			DeletedAt: &now,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}
