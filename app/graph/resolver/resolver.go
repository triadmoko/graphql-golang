package resolver

import (
	userService "github.com/triadmoko/grahpql-golang/service/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	User userService.UserServices
}
