package comment

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/repository/comment"
	"github.com/triadmoko/grahpql-golang/repository/post"
	"github.com/triadmoko/grahpql-golang/repository/user"
	"github.com/triadmoko/grahpql-golang/service/mail"
)

type CommentServices interface {
	Create(ctx context.Context, req request.NewComment) (*request.Comment, error)
	Update(ctx context.Context, id string, req request.NewComment) (*request.Comment, error)
	CommentList(ctx context.Context, req request.FilterComment) (*request.CommentList, error)
	Delete(ctx context.Context, id string) error
}
type comment_service struct {
	loggger           *logrus.Logger
	commentRepository comment.CommentRepository
	postRepository    post.PostRepository
	mailService       mail.MailServices
	userRepository    user.UserRepository
}

func NewCommentServices(
	loggger *logrus.Logger,
	commentRepository comment.CommentRepository,
	postRepository post.PostRepository,
	mailService mail.MailServices,
	userRepository user.UserRepository,
) *comment_service {
	return &comment_service{
		loggger:           loggger,
		commentRepository: commentRepository,
		postRepository:    postRepository,
		mailService:       mailService,
		userRepository:    userRepository,
	}
}
