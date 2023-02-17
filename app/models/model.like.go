package models

import (
	"time"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
)

type (
	Like struct {
		ID        string     `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
		UserID    string     `json:"user_id"`
		PostID    string     `json:"post_id"`
	}
)

func FormatterRequestLike(request request.NewLike) Like {
	response := Like{
		ID:        helpers.UUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		UserID:    *request.UserID,
		PostID:    request.PostID,
	}
	return response
}
func FormatterResponseLike(result Like) request.Like {
	response := request.Like{
		ID: result.ID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return response
}
