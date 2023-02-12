package post

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (s *post_service) Create(ctx context.Context, req request.NewPost) (*request.Post, error) {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return nil, errors.New("session invalid")
	}
	requestModel := models.FormatterRequestPost(sess.ID, req)
	resultPost, err := s.postRepository.Create(ctx, requestModel)
	if err != nil {
		s.loggger.Error("Error create post")
		return nil, errors.New("server not response")
	}
	response := models.FormatterResponsePost(*resultPost)
	return &response, nil
}
