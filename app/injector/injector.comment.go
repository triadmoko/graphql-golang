package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	commentRepository "github.com/triadmoko/grahpql-golang/repository/comment"
	postRepository "github.com/triadmoko/grahpql-golang/repository/post"
	userRepository "github.com/triadmoko/grahpql-golang/repository/user"
	commentService "github.com/triadmoko/grahpql-golang/service/comment"
)

func NewInjectorComment(conf *config.Config) commentService.CommentServices {
	repository := commentRepository.NewUserRepositorys(conf)
	repositoryPost := postRepository.NewPostRepositorys(conf)
	mailService := NewInjectorMail(conf)
	userRepository := userRepository.NewUserRepositorys(conf)
	service := commentService.NewCommentServices(conf.Logger, repository, repositoryPost, mailService, userRepository)
	return service
}
