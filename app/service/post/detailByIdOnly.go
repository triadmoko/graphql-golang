package post

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *post_service) DetailByIdOnly(ctx context.Context, id string) (*request.Post, error) {

	post, err := s.postRepository.DetailByIdOnly(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Post not found")
	}
	if err != nil {
		s.loggger.Error("Error get detail post")
		return nil, errors.New("server not response")
	}

	response := models.FormatterResponsePost(*post)
	return &response, nil
}
