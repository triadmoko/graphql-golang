package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	commentRepository "github.com/triadmoko/grahpql-golang/repository/comment"
	commentService "github.com/triadmoko/grahpql-golang/service/comment"
)

func NewInjectorComment(conf *config.Config) commentService.CommentServices {
	repository := commentRepository.NewUserRepositorys(conf)
	service := commentService.NewCommentServices(conf.Logger, repository)
	return service
}
