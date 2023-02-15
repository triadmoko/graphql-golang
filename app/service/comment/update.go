package comment

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *comment_service) Update(ctx context.Context, id string, req request.NewComment) (*request.Comment, error) {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return nil, errors.New("session invalid")
	}
	comment, err := s.commentRepository.Detail(ctx, id, sess.ID)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("comment not found")
	}
	if err != nil {
		s.loggger.Error("Error get detail comment")
		return nil, errors.New("server not response")
	}

	if comment == nil {
		return nil, errors.New("comment not found")
	}

	formatterReq := models.FormatterUpdateComment(req)
	comment, err = s.commentRepository.Update(ctx, id, formatterReq)
	if err != nil {
		s.loggger.Error("Error update comment")
		return nil, errors.New("server not response")
	}
	comment.ID = id
	response := models.FormatterResponseComment(*comment)
	return &response, nil
}
