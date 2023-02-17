package post

import (
	"context"

	"github.com/triadmoko/grahpql-golang/models"
)

func (r *post_repository) List(ctx context.Context, post models.PostList) (*models.PostList, error) {
	var posts []*models.Post
	where := "deleted_at IS NULL "
	var args []interface{}

	if len(post.Title) > 0 {
		where += "AND title like ?"
		args = append(args, "%"+post.Title+"%")
	}

	err := r.gorm.Scopes(models.PaginatePost(posts, &post, r.gorm, where, args...)).Where(where, args...).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	post.Posts = posts
	return &post, nil
}
