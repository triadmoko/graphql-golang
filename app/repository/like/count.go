package like

import (
	"context"
	"fmt"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (r *like_repository) Count(ctx context.Context, post_id string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(helpers.ContextTimeOut())*time.Second)
	defer cancel()

	var count int64
	err := r.gorm.Model(&models.Like{}).Where("deleted_at IS NULL AND post_id = ?", post_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	fmt.Println(count)
	return int(count), nil
}
