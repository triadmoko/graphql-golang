package user

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/model"
	"gorm.io/gorm"
)

func (s *user_service) Detail(ctx context.Context, id string) (*model.User, error) {
	result, err := s.userRepository.Detail(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("User not found")
	}
	if err != nil {
		s.loggger.Error("Error get detail user")
		return nil, errors.New("Server not response")
	}
	response := FormatterResponseUser(*result)
	return &response, nil
}
