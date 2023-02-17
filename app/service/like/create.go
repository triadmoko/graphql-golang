package like

import (
	"context"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *like_service) Like(ctx context.Context, req request.NewLike) (request.Like, error) {
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
