package user

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/gorm"
)

func (s *user_service) Login(ctx context.Context, req request.NewLogin) (*request.Token, error) {
	user, err := s.userRepository.GetOneByEmail(ctx, req.Email)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}

	if err != nil {
		s.loggger.Error("Failed Get User By Email ", err.Error())
		return nil, errors.New("server not response")
	}

	if user.Status == "inactive" {
		return nil, errors.New("user inactive, please check your email")
	}

	if err := helpers.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New("password is incorrect")
	}
	accessToken := map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
	}

	token, err := helpers.Sign(accessToken, "JWT_SECRET", 24*60*7)
	if err != nil {
		s.loggger.Error("invalid sign jwt", err.Error())
		return nil, errors.New("server not response")
	}

	return &request.Token{Token: token}, nil
}
