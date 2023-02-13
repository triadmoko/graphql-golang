package models

import (
	"time"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
)

type (
	Post struct {
		ID          string     `json:"id"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
		UserID      string     `json:"user_id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
	}
)

func FormatterRequestPost(userID string, request request.NewPost) Post {
	response := Post{
		ID:          helpers.UUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedAt:   nil,
		UserID:      userID,
		Title:       request.Title,
		Description: request.Description,
	}
	return response
}

func FormatterResponsePost(post Post) request.Post {
	response := request.Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Description: post.Description,
	}
	return response
}

func FormatterUpdatePost(post request.NewPost) Post {
	response := Post{
		Title:       post.Title,
		Description: post.Description,
		UpdatedAt:   time.Now().UTC(),
	}
	return response
}
