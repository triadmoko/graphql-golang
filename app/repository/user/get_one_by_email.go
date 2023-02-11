package user

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *user_repository) GetOneByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	user := models.User{}
	err := r.gorm.First(&user, "deleted_at IS NULL AND email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
