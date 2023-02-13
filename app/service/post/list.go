package post

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/models"
)

func (s *post_service) List(ctx context.Context, req request.FilterPost) (*request.PostList, error) {
	filter := models.FilterPost(req)
	posts, err := s.postRepository.List(ctx, filter)
	if err != nil {
		s.loggger.Error("Error get post list")
		return nil, errors.New("server not response")
	}
	response := models.FormatterResponsePostList(*posts)
	return &response, nil
}
