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
	comment, err := s.commentRepository.Detail(ctx, id, sess.ID)
	if err.Error() == gorm.ErrRecordNotFound.Error() {
		return errors.New("Comment not found")
	}
	if err != nil {
		s.loggger.Error("Error get detail comment")
		return errors.New("server not response")
	}
	err = s.commentRepository.Delete(ctx, comment.ID, comment.UserID)
	if err != nil {
		s.loggger.Error("Error delete comment")
		return errors.New("server not response")
	}

	return nil
}
