package mail

import (
	"context"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *mail_repository) Create(ctx context.Context, request models.EmailVerification) (*models.EmailVerification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	err := r.gorm.WithContext(ctx).Model(&models.EmailVerification{}).Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
