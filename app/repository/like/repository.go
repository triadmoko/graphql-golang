package like

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

type LikeRepository interface {
	Create(ctx context.Context, request models.Like) (*models.Like, error)
	Delete(ctx context.Context, id string) error
	Detail(ctx context.Context, userID, postID string) (*models.Like, error)
	Count(ctx context.Context, post_id string) (int, error)
}
type like_repository struct {
	loggger *logrus.Logger
	gorm    *gorm.DB
}

func NewLikeRepositorys(config *config.Config) *like_repository {
	return &like_repository{loggger: config.Logger, gorm: config.Gorm}
}
