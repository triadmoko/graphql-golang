package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	postService "github.com/triadmoko/grahpql-golang/service/post"
	userService "github.com/triadmoko/grahpql-golang/service/user"
)

type Injector struct {
	User userService.UserServices
	Post postService.PostServices
}

func NewInitInjector(conf *config.Config) *Injector {
	return &Injector{
		User: NewInjectorUser(conf),
		Post: NewInjectorPost(conf),
	}
}
