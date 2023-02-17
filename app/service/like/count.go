package like

import "context"

func (s *like_service) Count(ctx context.Context, postID string) (int, error) {
	return s.likeRepository.Count(ctx, postID)
}
