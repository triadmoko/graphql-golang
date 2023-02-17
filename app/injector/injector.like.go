package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	likeRepository "github.com/triadmoko/grahpql-golang/repository/like"
	likeService "github.com/triadmoko/grahpql-golang/service/like"
)

func NewInjectorLike(conf *config.Config) likeService.LikeService {
	repository := likeRepository.NewLikeRepositorys(conf)
	service := likeService.NewLikeServices(conf.Logger, repository)
	return service
}
