package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	userService "github.com/triadmoko/grahpql-golang/service/user"
)

type Injector struct {
	User userService.UserServices
}

func NewInitInjector(conf *config.Config) *Injector {
	return &Injector{
		User: NewInjectorUser(conf),
	}
}
