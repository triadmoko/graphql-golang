package like

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/repository/like"
)

type LikeService interface {
	Like(ctx context.Context, req request.NewLike) (request.Like, error)
	Count(ctx context.Context, postID string) (int, error)
}
type like_service struct {
	loggger        *logrus.Logger
	likeRepository like.LikeRepository
}

func NewLikeServices(
	loggger *logrus.Logger,
	likeRepository like.LikeRepository,
) *like_service {
	return &like_service{
		loggger:        loggger,
		likeRepository: likeRepository,
	}
}
