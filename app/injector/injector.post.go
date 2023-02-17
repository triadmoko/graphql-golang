package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	postRepository "github.com/triadmoko/grahpql-golang/repository/post"
	userRepository "github.com/triadmoko/grahpql-golang/repository/user"
	postService "github.com/triadmoko/grahpql-golang/service/post"
)

func NewInjectorPost(conf *config.Config) postService.PostServices {
	repository := postRepository.NewPostRepositorys(conf)
	userRepository := userRepository.NewUserRepositorys(conf)
	likeService := NewInjectorLike(conf)
	service := postService.NewPostServices(conf.Logger, repository, userRepository, likeService)
	return service
}
