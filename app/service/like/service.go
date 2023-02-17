package like

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/repository/like"
	"github.com/triadmoko/grahpql-golang/repository/post"
	"github.com/triadmoko/grahpql-golang/repository/user"
	"github.com/triadmoko/grahpql-golang/service/mail"
)

type LikeService interface {
	Like(ctx context.Context, req request.NewLike) (request.Like, error)
	Count(ctx context.Context, postID string) (int, error)
}
type like_service struct {
	loggger        *logrus.Logger
	likeRepository like.LikeRepository
	postRepository post.PostRepository
	mailService    mail.MailServices
	userRepository user.UserRepository
}

func NewLikeServices(
	loggger *logrus.Logger,
	likeRepository like.LikeRepository,
	postRepository post.PostRepository,
	mailService mail.MailServices,
	userRepository user.UserRepository,
) *like_service {
	return &like_service{
		loggger:        loggger,
		likeRepository: likeRepository,
		postRepository: postRepository,
		mailService:    mailService,
		userRepository: userRepository,
	}
}
