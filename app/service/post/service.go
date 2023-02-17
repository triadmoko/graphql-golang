package post

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/repository/post"
	"github.com/triadmoko/grahpql-golang/repository/user"
	"github.com/triadmoko/grahpql-golang/service/like"
)

type PostServices interface {
	Create(ctx context.Context, req request.NewPost) (*request.Post, error)
	Update(ctx context.Context, id string, req request.NewPost) (*request.Post, error)
	Detail(ctx context.Context, id string) (*request.Post, error)
	DetailByIdOnly(ctx context.Context, id string) (*request.Post, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, req request.FilterPost) (*request.PostList, error)
}
type post_service struct {
	loggger        *logrus.Logger
	postRepository post.PostRepository
	userRepository user.UserRepository
	likeService    like.LikeService
}

func NewPostServices(
	loggger *logrus.Logger,
	postRepository post.PostRepository,
	userRepository user.UserRepository,
	likeService like.LikeService,
) *post_service {
	return &post_service{
		loggger:        loggger,
		postRepository: postRepository,
		userRepository: userRepository,
		likeService:    likeService,
	}
}
