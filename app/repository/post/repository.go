package post

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(ctx context.Context, request models.Post) (*models.Post, error)
	Update(ctx context.Context, id string, post models.Post) (*models.Post, error)
	Detail(ctx context.Context, id, userID string) (*models.Post, error)
	Delete(ctx context.Context, id, userID string) error
	List(ctx context.Context, post models.PostList) (*models.PostList, error)
}
type post_repository struct {
	loggger *logrus.Logger
	gorm    *gorm.DB
}

func NewPostRepositorys(config *config.Config) *post_repository {
	return &post_repository{loggger: config.Logger, gorm: config.Gorm}
}
