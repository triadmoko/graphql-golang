package post

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *post_service) Update(ctx context.Context, id string, req request.NewPost) (*request.Post, error) {
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

	if post == nil {
		return nil, errors.New("Post not found")
	}

	formatterReq := models.FormatterUpdatePost(req)
	post, err = s.postRepository.Update(ctx, id, formatterReq)
	if err != nil {
		s.loggger.Error("Error update post")
		return nil, errors.New("server not response")
	}
	post.ID = id
	response := models.FormatterResponsePost(*post)
	return &response, nil
}
