package models

import (
	"math"
	"time"

	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
	"gorm.io/gorm"
)

type (
	Comment struct {
		ID          string     `json:"id"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
		UserID      string     `json:"user_id"`
		PostID      string     `json:"post_id"`
		Description string     `json:"description"`
	}
	CommentList struct {
		PostID    string     `json:"post_id,omitempty"`
		PerPage   int        `json:"per_page,omitempty;query:limit"`
		Page      int        `json:"page,omitempty;query:page"`
		Sort      string     `json:"sort,omitempty;query:sort"`
		TotalData int64      `json:"total_data"`
		TotalPage int        `json:"total_page"`
		Comments  []*Comment `json:"comments"`
	}
)

func (p *CommentList) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *CommentList) GetLimit() int {
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	return p.PerPage
}

func (p *CommentList) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *CommentList) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func PaginateComment(value interface{}, pagination *CommentList, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalData = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.PerPage)))
	pagination.TotalPage = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func FormatterRequestComment(userID string, request request.NewComment) Comment {
	response := Comment{
		ID:          helpers.UUID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedAt:   nil,
		PostID:      request.PostID,
		UserID:      userID,
		Description: request.Description,
	}
	return response
}

func FormatterResponseComment(post Comment) request.Comment {
	response := request.Comment{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Description: post.Description,
	}
	return response
}

func FormatterUpdateComment(post request.NewComment) Comment {
	response := Comment{
		Description: post.Description,
		UpdatedAt:   time.Now().UTC(),
	}
	return response
}

func FilterComment(filter request.FilterComment) CommentList {
	response := CommentList{
		PostID:  filter.PostID,
		PerPage: filter.PerPage,
		Page:    filter.Page,
	}
	return response
}

func FormatterResponseCommentList(results CommentList) request.CommentList {
	comments := request.CommentList{
		Pagination: &request.Pagination{
			PerPage:   results.PerPage,
			Page:      results.Page,
			TotalPage: results.TotalPage,
			TotalData: int(results.TotalData),
		},
	}

	for _, comment := range results.Comments {
		formatterComment := FormatterResponseComment(*comment)
		comments.Comments = append(comments.Comments, &formatterComment)
	}

	return comments
}
