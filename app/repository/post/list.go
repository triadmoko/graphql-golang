package post

import (
	"context"

	"github.com/triadmoko/grahpql-golang/models"
)

func (r *post_repository) List(ctx context.Context, post models.PostList) (*models.PostList, error) {
	var posts []*models.Post
	err := r.gorm.Scopes(models.Paginate(posts, &post, r.gorm)).Where("deleted_at IS NULL").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	post.Posts = posts
	return &post, nil
}
