package post

import (
	"context"
	"errors"
	"fmt"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (s *post_service) Create(ctx context.Context, req request.NewPost) (*request.Post, error) {
	fmt.Println("1")
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return nil, errors.New("session invalid")
	}
	if sess.ID == "" || len(sess.ID) == 0 {
		return nil, errors.New("request login")
	}
	resultPost, err := s.postRepository.Create(ctx, models.FormatterRequestPost(sess.ID, req))
	if err != nil {
		s.loggger.Error("Error create post")
		return nil, errors.New("server not response")
	}

	resultUser, err := s.userService.Detail(ctx, sess.ID)
	if err != nil {
		s.loggger.Error("Error get user by id")
		return nil, errors.New("server not response")
	}
	response := models.FormatterResponsePost(*resultPost, *resultUser)
	return &response, nil
}
