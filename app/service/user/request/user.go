package request

import (
	"time"

	"github.com/triadmoko/grahpql-golang/graph/model"
	"github.com/triadmoko/grahpql-golang/helpers"
)

func FormmaterRequestUser(request model.NewUser) model.User {
	response := model.User{
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
