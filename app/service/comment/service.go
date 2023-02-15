package comment

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/repository/comment"
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
}

func NewCommentServices(
	loggger *logrus.Logger,
	commentRepository comment.CommentRepository,
) *comment_service {
	return &comment_service{
		loggger:           loggger,
		commentRepository: commentRepository,
	}
}