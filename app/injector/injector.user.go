package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	userRepository "github.com/triadmoko/grahpql-golang/repository/user"
	userService "github.com/triadmoko/grahpql-golang/service/user"
)

func NewInjectorUser(conf *config.Config) userService.UserServices {
	mail := NewInjectorMail(conf)
	repository := userRepository.NewUserRepositorys(conf)
	service := userService.NewUserServices(conf.Logger, repository, mail)
	return service
}
