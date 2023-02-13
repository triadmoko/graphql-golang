package post

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (s *post_service) Detail(ctx context.Context, id string) (*request.Post, error) {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return nil, errors.New("session invalid")
	}
	post, err := s.postRepository.Detail(ctx, id, sess.ID)
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
