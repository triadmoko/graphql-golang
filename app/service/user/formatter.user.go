package user

import (
	"time"

	"github.com/triadmoko/grahpql-golang/graph/model"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func FormatterRequestUser(request model.NewUser) models.User {
	response := models.User{
		ID:        helpers.UUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Status:    "inactive",
	}
	return response
}

func FormatterResponseUser(result models.User) model.User {
	response := model.User{
		ID:        result.ID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		Name:      result.Name,
		Email:     result.Email,
		Status:    result.Status,
	}
	return response
}
