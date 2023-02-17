package comment

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/gorm"
)

func (s *comment_service) Delete(ctx context.Context, id string) error {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return errors.New("session invalid")
	}
	comment, err := s.commentRepository.Detail(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return errors.New("Comment not found")
	}
	if err != nil {
		s.loggger.Error("Error get detail comment")
		return errors.New("server not response")
	}

	post, err := s.postRepository.DetailByIdOnly(ctx, comment.PostID)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("Error get detail comment")
		return errors.New("server not response")
	}

	if post.UserID == sess.ID { // is owner post
		if err := s.commentRepository.Delete(ctx, comment.ID, comment.UserID); err != nil {
			s.loggger.Error("Error delete comment")
			return errors.New("server not response")
		}
	} else if comment.UserID == sess.ID { // is owner comment
		if err := s.commentRepository.Delete(ctx, comment.ID, comment.UserID); err != nil {
			s.loggger.Error("Error delete comment")
			return errors.New("server not response")
		}
	} else {
		return errors.New("you are not allow delete comment")
	}
	return nil
}
