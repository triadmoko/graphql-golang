package models

import (
	"time"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
)

type (
	User struct {
		ID        string     `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		Status    string     `json:"status"`
		Password  string     `json:"password"`
	}
)

func FormatterRequestUser(request request.NewUser) User {
	response := User{
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

func FormatterRequestUpdateUser(request request.UpdateUser, status string) User {
	response := User{
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		Name:      request.Name,
		Email:     request.Email,
		Status:    status,
	}

	if request.Password != nil {
		response.Password = *request.Password
	}

	return response
}

func FormatterResponseUser(result User) request.User {
	response := request.User{
		ID:        result.ID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		Name:      result.Name,
		Email:     result.Email,
		Status:    result.Status,
	}
	return response
}
