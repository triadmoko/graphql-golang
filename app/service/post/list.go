package post

import (
	"context"
	"errors"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func (s *post_service) List(ctx context.Context, req request.FilterPost) (*request.PostList, error) {
	filter := models.FilterPost(req)
	posts, err := s.postRepository.List(ctx, filter)
	if err != nil {
		s.loggger.Error("Error get post list")
		return nil, errors.New("server not response")
	}
	preloads := helpers.GetPreloads(ctx)
	response := models.FormatterResponsePostList(*posts)

	for _, preload := range preloads {
		switch preload {
		case "posts.user":
			for i, post := range posts.Posts {
				result, err := s.userRepository.Detail(ctx, post.UserID)
				if err != nil {
					s.loggger.Error("Error get user detail")
					return nil, errors.New("server not response")
				}
				user := models.FormatterResponseUser(*result)
				response.Posts[i].User = &user
			}
		case "posts.total_like":
			for i, post := range posts.Posts {
				count, err := s.likeService.Count(ctx, post.ID)
				if err != nil {
					s.loggger.Error("Error get user detail")
					return nil, errors.New("server not response")
				}
				response.Posts[i].TotalLike = &count
			}
		}

	}
	return &response, nil
}
