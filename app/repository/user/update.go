package user

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *user_repository) Update(ctx context.Context, id string, user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.User{}).Where("deleted_at IS NULL AND id = ?", id).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
