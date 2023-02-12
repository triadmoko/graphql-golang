package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	postRepository "github.com/triadmoko/grahpql-golang/repository/post"
	postService "github.com/triadmoko/grahpql-golang/service/post"
)

func NewInjectorPost(conf *config.Config) postService.PostServices {
	repository := postRepository.NewPostRepositorys(conf)
	service := postService.NewPostServices(conf.Logger, repository)
	return service
}
