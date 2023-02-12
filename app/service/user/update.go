package user

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"gorm.io/gorm"
)

func (s *user_service) Update(ctx context.Context, id string, req *request.UpdateUser) (*request.User, error) {
	user, err := s.userRepository.Detail(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("User not found")
	}

	if err != nil {
		s.loggger.Error("Error get detail user ", err.Error())
		return nil, errors.New("Server not response")
	}

	if req.Password != nil {
		password, err := helpers.HashPassword(*req.Password)
		if err != nil {
			s.loggger.Error("Error HashPassword ", err.Error())
			return nil, errors.New("Failed Hass Password")
		}
		req.Password = &password
	}

	request := models.FormatterRequestUpdateUser(*req, user.Status)
	result, err := s.userRepository.Update(ctx, id, request)
	if err != nil {
		s.loggger.Error("Failed Update User ", err.Error())
		return nil, errors.New("server not response")
	}

	response := models.FormatterResponseUser(*result)
	return &response, nil
}
