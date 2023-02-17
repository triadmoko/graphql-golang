package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	commentService "github.com/triadmoko/grahpql-golang/service/comment"
	likeService "github.com/triadmoko/grahpql-golang/service/like"
	postService "github.com/triadmoko/grahpql-golang/service/post"
	userService "github.com/triadmoko/grahpql-golang/service/user"
)

type Injector struct {
	User    userService.UserServices
	Post    postService.PostServices
	Comment commentService.CommentServices
	Like    likeService.LikeService
}

func NewInitInjector(conf *config.Config) *Injector {
	return &Injector{
		User:    NewInjectorUser(conf),
		Post:    NewInjectorPost(conf),
		Comment: NewInjectorComment(conf),
		Like:    NewInjectorLike(conf),
	}
}
