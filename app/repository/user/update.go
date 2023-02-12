package user

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *user_repository) UpdateStatusUser(ctx context.Context, id, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.User{}).Where("deleted_at IS NULL AND id = ?", id).Updates(&models.User{
		Status: status,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
