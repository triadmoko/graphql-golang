package user

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/model"
	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/gorm"
)

func (s *user_service) Create(ctx context.Context, req model.NewUser) (*model.User, error) {
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

	request := FormatterRequestUser(req)
	result, err := s.userRepository.Create(ctx, request)
	if err != nil {
		s.loggger.Error("Failed Create User ", err.Error())
		return nil, errors.New("server not response")
	}

	response := FormatterResponseUser(*result)
	return &response, nil
}
