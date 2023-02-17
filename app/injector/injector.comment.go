package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	commentRepository "github.com/triadmoko/grahpql-golang/repository/comment"
	postRepository "github.com/triadmoko/grahpql-golang/repository/post"
	commentService "github.com/triadmoko/grahpql-golang/service/comment"
)

func NewInjectorComment(conf *config.Config) commentService.CommentServices {
	repository := commentRepository.NewUserRepositorys(conf)
	repositoryPost := postRepository.NewPostRepositorys(conf)
	service := commentService.NewCommentServices(conf.Logger, repository, repositoryPost)
	return service
}
