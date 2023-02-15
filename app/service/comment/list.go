package comment

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/models"
)

func (s *comment_service) CommentList(ctx context.Context, req request.FilterComment) (*request.CommentList, error) {
	filter := models.FilterComment(req)
	posts, err := s.commentRepository.List(ctx, filter)
	if err != nil {
		s.loggger.Error("Error get comment list")
		return nil, errors.New("server not response")
	}
	response := models.FormatterResponseCommentList(*posts)
	return &response, nil
}
