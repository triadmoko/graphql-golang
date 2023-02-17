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
	post, err := s.postRepository.DetailByIdOnly(ctx, req.PostID)
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
	user, err := s.userRepository.Detail(ctx, post.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("Error get detail User")
		return nil, errors.New("server not response")
	}
	
	response := models.FormatterResponseComment(*resultPost)
	
	if post.UserID != sess.ID {
		go s.mailService.SendEmailCommentAndLike(
			models.FormSendEmail{
				Subject: "Comment",
				Text:    "Postingan di komentar oleh " + sess.Name,
				To:      []string{user.Email},
			},
		)
	}

	return &response, nil
}
