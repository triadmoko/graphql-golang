package comment

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *comment_service) Create(ctx context.Context, req request.NewComment) (*request.Comment, error) {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return nil, errors.New("session invalid")
	}
	_, err := s.postRepository.DetailByIdOnly(ctx, req.PostID)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Post Not Found")
	}
	
	if err != nil {
		return nil, errors.New("server not response")
	}

	requestModel := models.FormatterRequestComment(sess.ID, req)
	resultPost, err := s.commentRepository.Create(ctx, requestModel)
	if err != nil {
		s.loggger.Error("Error create comment")
		return nil, errors.New("server not response")
	}
	response := models.FormatterResponseComment(*resultPost)
	return &response, nil
}
