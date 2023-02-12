package user

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *user_service) Create(ctx context.Context, req request.NewUser) (*request.User, error) {
	user, err := s.userRepository.GetOneByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.loggger.Error("Failed Get User By Email ", err.Error())
		return nil, errors.New("server not response")
	}

	if user != nil {
		return nil, errors.New("User Already Register")
	}

	req.Password, err = helpers.HashPassword(req.Password)
	if err != nil {
		s.loggger.Error("Error HashPassword ", err.Error())
		return nil, errors.New("Failed Hass Password")
	}

	request := models.FormatterRequestUser(req)
	result, err := s.userRepository.Create(ctx, request)
	if err != nil {
		s.loggger.Error("Failed Create User ", err.Error())
		return nil, errors.New("server not response")
	}

	err = s.mailService.EmailVerification(ctx, result.Email, result.ID)
	if err != nil {
		s.loggger.Error("Failed Create User ", err.Error())
		return nil, err
	}

	response := models.FormatterResponseUser(*result)
	return &response, nil
}
