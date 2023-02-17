package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	likeRepository "github.com/triadmoko/grahpql-golang/repository/like"
	postRepository "github.com/triadmoko/grahpql-golang/repository/post"
	userRepository "github.com/triadmoko/grahpql-golang/repository/user"
	likeService "github.com/triadmoko/grahpql-golang/service/like"
)

func NewInjectorLike(conf *config.Config) likeService.LikeService {
	repository := likeRepository.NewLikeRepositorys(conf)
	repositoryPost := postRepository.NewPostRepositorys(conf)
	mailService := NewInjectorMail(conf)
	userRepository := userRepository.NewUserRepositorys(conf)
	service := likeService.NewLikeServices(conf.Logger, repository, repositoryPost, mailService, userRepository)
	return service
}
