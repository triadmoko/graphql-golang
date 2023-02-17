// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package request

import (
	"time"
)

type Comment struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Description string     `json:"description"`
	User        *User      `json:"user"`
	Post        *Post      `json:"post"`
	Like        []*Like    `json:"like"`
}

type CommentList struct {
	Comments   []*Comment  `json:"comments"`
	Pagination *Pagination `json:"pagination"`
}

type FilterComment struct {
	PostID  string `json:"post_id"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

type FilterPost struct {
	ID      []*string `json:"id"`
	Title   *string   `json:"title"`
	UserID  []*string `json:"user_id"`
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
}

type Like struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Post      *Post      `json:"post"`
	User      *User      `json:"user"`
}

type NewComment struct {
	UserID      *string `json:"user_id"`
	PostID      string  `json:"post_id"`
	Description string  `json:"description"`
}

type NewLike struct {
	UserID *string `json:"user_id"`
	PostID string  `json:"post_id"`
}

type NewLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewVerify struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

type Pagination struct {
	PerPage   int `json:"per_page"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

type Post struct {
	ID          string       `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	TotalLike   *int         `json:"total_like"`
	User        *User        `json:"user"`
	Comments    *CommentList `json:"comments"`
}

type PostList struct {
	Posts      []*Post     `json:"posts"`
	Pagination *Pagination `json:"pagination"`
}

type Token struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
}
