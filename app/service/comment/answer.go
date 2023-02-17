package comment

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *comment_service) Answer(ctx context.Context, commentID string, req request.NewAnswerComment) (*request.Comment, error) {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return nil, errors.New("session invalid")
	}
	comment, err := s.commentRepository.Detail(ctx, commentID)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Comment Not Found")
	}

	if err != nil {
		return nil, errors.New("server not response")
	}

	requestModel := models.FormatterRequestAnswerComment(sess.ID, comment.ID, comment.PostID, req)
	resultPost, err := s.commentRepository.Create(ctx, requestModel)
	if err != nil {
		s.loggger.Error("Error create comment")
		return nil, errors.New("server not response")
	}
	user, err := s.userRepository.Detail(ctx, comment.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("Error get detail User")
		return nil, errors.New("server not response")
	}
	resultPost.PostID = comment.PostID
	response := models.FormatterResponseComment(*resultPost)

	if comment.UserID != sess.ID {
		go s.mailService.SendEmailCommentAndLike(
			models.FormSendEmail{
				Subject: "Answer Comment",
				Text:    "Komentar Anda di komentari oleh " + sess.Name,
				To:      []string{user.Email},
			},
		)
	}

	return &response, nil
}
