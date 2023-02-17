package models

import (
	"math"
	"time"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/gorm"
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
	PostList struct {
		Title     string  `json:"title,omitempty"`
		PerPage   int     `json:"per_page,omitempty;query:limit"`
		Page      int     `json:"page,omitempty;query:page"`
		Sort      string  `json:"sort,omitempty;query:sort"`
		TotalData int64   `json:"total_data"`
		TotalPage int     `json:"total_page"`
		Posts     []*Post `json:"posts"`
	}
)

func (p *PostList) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PostList) GetLimit() int {
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	return p.PerPage
}

func (p *PostList) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PostList) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func PaginatePost(value interface{}, pagination *PostList, db *gorm.DB, where string , args ...interface{}) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(&Post{}).Where(where, args...).Count(&totalRows)

	pagination.TotalData = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.PerPage)))
	pagination.TotalPage = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

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

func FilterPost(filter request.FilterPost) PostList {
	response := PostList{
		PerPage: filter.PerPage,
		Page:    filter.Page,
	}
	if filter.Title != nil {
		response.Title = *filter.Title
	}
	return response
}

func FormatterResponsePostList(results PostList) request.PostList {
	posts := request.PostList{
		Pagination: &request.Pagination{
			PerPage:   results.PerPage,
			Page:      results.Page,
			TotalPage: results.TotalPage,
			TotalData: int(results.TotalData),
		},
	}

	for _, post := range results.Posts {
		post := FormatterResponsePost(*post)
		posts.Posts = append(posts.Posts, &post)
	}

	return posts
}
