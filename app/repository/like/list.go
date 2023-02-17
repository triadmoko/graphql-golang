package like

import (
	"context"

	"github.com/triadmoko/grahpql-golang/models"
)

func (r *like_repository) List(ctx context.Context, comment models.CommentList) (*models.CommentList, error) {
	var comments []*models.Comment
	err := r.gorm.Scopes(models.PaginateComment(comments, &comment, r.gorm)).
		Where("deleted_at IS NULL AND post_id = ?", comment.PostID).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	comment.Comments = comments
	return &comment, nil
}
