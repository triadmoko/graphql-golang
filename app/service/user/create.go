package user

import (
	"github.com/triadmoko/grahpql-golang/graph/model"
	"github.com/triadmoko/grahpql-golang/service/user/request"
)

func (s *user_service) Create(req model.NewUser) (model.User, error) {
	
	return request.FormmaterRequestUser(req), nil
}
