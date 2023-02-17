package comment

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, request models.Comment) (*models.Comment, error)
	Delete(ctx context.Context, id, userID string) error
	Update(ctx context.Context, id string, comment models.Comment) (*models.Comment, error)
	List(ctx context.Context, comment models.CommentList) (*models.CommentList, error)
	Detail(ctx context.Context, id string) (*models.Comment, error)
}
type comment_repository struct {
	loggger *logrus.Logger
	gorm    *gorm.DB
}

func NewUserRepositorys(config *config.Config) *comment_repository {
	return &comment_repository{loggger: config.Logger, gorm: config.Gorm}
}
