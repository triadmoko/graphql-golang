package post

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/gorm"
)

func (s *post_service) Delete(ctx context.Context, id string) error {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return errors.New("session invalid")
	}
	post, err := s.postRepository.Detail(ctx, id, sess.ID)
	if err == gorm.ErrRecordNotFound {
		return errors.New("Post not found")
	}
	if err != nil {
		s.loggger.Error("Error get detail post")
		return errors.New("server not response")
	}
	err = s.postRepository.Delete(ctx, post.ID, post.UserID)
	if err != nil {
		s.loggger.Error("Error delete post")
		return errors.New("server not response")
	}

	return nil
}
