package user

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *user_repository) GetVerifyByEmail(ctx context.Context, email string) (*models.EmailVerification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()
	verify := models.EmailVerification{}
	err := r.gorm.First(&verify, "deleted_at IS NULL AND email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &verify, nil
}
