package like

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *like_service) Like(ctx context.Context, req request.NewLike) (request.Like, error) {
	sess, ok := ctx.Value("sess").(*helpers.MetaToken)
	if !ok {
		return request.Like{}, errors.New("session invalid")
	}
	post, err := s.postRepository.DetailByIdOnly(ctx, req.PostID)
	if err == gorm.ErrRecordNotFound {
		return request.Like{}, errors.New("Post Not Found")
	}
	if err != nil {
		s.loggger.Error("Error get detail post by id only")
		return request.Like{}, nil
	}
	user, err := s.userRepository.Detail(ctx, post.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("Error get detail post by id only")
		return request.Like{}, nil
	}
	var isNewLike bool
	like, err := s.likeRepository.Detail(ctx, *req.UserID, req.PostID)
	if err == gorm.ErrRecordNotFound {
		isNewLike = true
	} else {
		isNewLike = false
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("Error Like Post")
		return request.Like{}, err
	}
	if isNewLike {
		mapLike := models.FormatterRequestLike(req)
		result, err := s.likeRepository.Create(ctx, mapLike)
		if err != nil {
			s.loggger.Error("Error Like Post")
			return request.Like{}, err
		}
		mapResponse := models.FormatterResponseLike(*result)
		go s.mailService.SendEmailCommentAndLike(
			models.FormSendEmail{
				Subject: "Like",
				Text:    "Postingan kamu di sukai oleh " + sess.Name,
				To:      []string{user.Email},
			},
		)
		
		return mapResponse, nil
	} else {
		err := s.likeRepository.Delete(ctx, like.ID)
		if err != nil {
			s.loggger.Error("Error Like Post")
			return request.Like{}, err
		}
		mapResponse := models.FormatterResponseLike(models.Like{ID: like.ID, CreatedAt: like.CreatedAt, UpdatedAt: like.UpdatedAt})
		return mapResponse, nil
	}
}
