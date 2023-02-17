package resolver

import (
	commentService "github.com/triadmoko/grahpql-golang/service/comment"
	likeService "github.com/triadmoko/grahpql-golang/service/like"
	postService "github.com/triadmoko/grahpql-golang/service/post"
	userService "github.com/triadmoko/grahpql-golang/service/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//
//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	User    userService.UserServices
	Post    postService.PostServices
	Comment commentService.CommentServices
	LikeService    likeService.LikeService
}
